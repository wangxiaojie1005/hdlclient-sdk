package operate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/wangxiaojie1005/hdlclient-sdk/model"
	//"hdlclient/model/httpModel"
	"github.com/wangxiaojie1005/hdlclient-sdk/units"
	"io"
	"net"
	"net/http"
	"net/url"
)

type Request struct {
}

// 定义要发送和接收的 JSON 数据的结构体
type RequestData struct {
	handle string `json:"message"`
}

type ResponseData struct {
	Reply string `json:"reply"`
}

func (req *Request) SendUdpMsg(msg []byte) (interface{}, error) {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		//IP: net.IPv4(61, 147, 200, 66),
		//IP:   net.IPv4(172, 17, 6, 29),219.134.89.12
		//Port: 7641,
		//IP:   net.IPv4(219, 134, 89, 12),
		//IP: net.IPv4(127, 0, 0, 1),
		//IP: net.IPv4(113, 209, 196, 34), //handle
		//IP: net.IPv4(36, 112, 25, 8), //BJ VAA
		//IP: net.IPv4(113, 209, 196, 35), //113.209.196.35  86.1000.16
		//IP: net.IPv4(124, 127, 41, 238), //124.127.41.238 86.1000.16/1440.9.011.0411-20161017
		//IP:   net.IPv4(221, 123, 181, 189),
		//IP:   net.IPv4(36, 112, 25, 8), //88.118
		IP:   net.IPv4(208, 254, 38, 90), //208.254.38.90   10.1016/j.elerap.2011.05.004
		Port: 2641,

		//IP:   net.IPv4(172, 18, 6, 40),
		//Port: 5641,
	})

	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return nil, err // 返回空的字节切片
	}
	defer socket.Close()

	sendData := msg
	_, err = socket.Write(sendData) // 发送数据
	if err != nil {
		fmt.Println("发送数据失败，err:", err)
		return nil, err // 返回空的字节切片
	}

	println("\n\n")
	f := 0

	//接收服务端的值
	data := make([]byte, 512)

	//socket.ReadFromUDP(data[:])
	//n:读取到的数据字节数 remoteAddr: 这个变量存储了发送数据的远程主机的地址
	n, remoteAddr, err := socket.ReadFromUDP(data[:]) // 接收数据
	if err != nil {
		fmt.Println("接收数据失败，err:", err)
		return nil, err // 返回空的字节切片
	}
	fmt.Printf("received firstPackage. addr:%v \ncount:%v\n", remoteAddr, n)
	fmt.Printf("data的类型，%T", data)
	header, rec := ProcFistPackage(data, "UDP") //判断是否有第二个或更多包
	obj := len(data[44:])
	//headerData := data[22:44]
	bodyData := data[44:] //把第一包的body数据保存到bodyData中
	fmt.Printf("bodyData的类型，%T", bodyData)
	fo := uint32(obj)
	fo = header.BodyLength - fo //还有包
	if rec {
		for {
			f = f + 1
			fmt.Printf(">>>>>>%d>>>>>>%d>>>>\n", f, fo)

			if fo >= 512 {
				fo = fo - 512
				dataCycle := make([]byte, 512)
				socket.ReadFromUDP(dataCycle[:])
				env := new(model.MessageEnvelope)
				env.ToObject(dataCycle[:20])
				bodyData = append(bodyData, dataCycle[20:]...)
			} else {
				//dataCycle := make([]byte, fo+24)
				dataCycle := make([]byte, 512)
				socket.ReadFromUDP(dataCycle[:])
				env := new(model.MessageEnvelope)
				env.ToObject(dataCycle[:20])
				bodyData = append(bodyData, dataCycle[20:]...)
				break
			}

		}

	}
	//fmt.Printf("rev packages count:%d\n  ", len(bodyData))
	fmt.Printf("bodyData的类型，%T", bodyData)
	if header.ResponseCode == units.RC_SUCCESS {
		if header.OpCode == units.OC_RESOLUTION {
			responseData := ResolveBodyData(bodyData, header.OpCode) //解析bodyData
			return responseData, nil
		} else {
			return nil, nil
		}
	} else {
		responseData := model.DecodeErrBodyFromBytes(bodyData) //解析bodyData
		return responseData, nil
	}
}
func (req *Request) SendTcpMsg(msg []byte) (interface{}, error) {
	// 连接服务器
	//conn, err := net.Dial("tcp", "36.112.25.14:2641")
	//conn, err := net.Dial("tcp", "208.254.38.90:2641")
	conn, err := net.Dial("tcp", "172.18.6.40:5641") // qth 测试IP
	if err != nil {
		fmt.Println("无法连接到服务器:", err)
		return nil, err // 返回空的字节切片
	}
	defer conn.Close()

	// 发送数据
	sendData := msg
	// 发送数据
	_, err = conn.Write(sendData)
	if err != nil {
		fmt.Println("无法发送数据到服务器:", err)
		return nil, err // 返回空的字节切片
	}

	// 接收数据
	recvData := make([]byte, 1024) // 假设接收缓冲区大小为1024字节
	_, err = conn.Read(recvData)
	if err != nil {
		fmt.Println("无法从服务器接收数据:", err)
		return nil, err
	}
	fmt.Println("从服务器收到的字节:", recvData)
	header, rec := ProcFistPackage(recvData, "TCP")

	fmt.Println("rec----------:", rec)
	bodyData := recvData[44:]
	if header.ResponseCode == units.RC_SUCCESS {
		if header.OpCode == units.OC_RESOLUTION {
			responseData := ResolveBodyData(bodyData, header.OpCode) //解析bodyData
			return responseData, nil
		} else {
			return nil, nil
		}
	} else {
		fmt.Println("Rec Error,ReponseCode:", header.ResponseCode)
		responseData := model.DecodeErrBodyFromBytes(bodyData) //解析bodyDat
		return responseData, nil
	}
}
func (req *Request) SendHttpMsg(msg string, method string, params interface{}) (interface{}, error) {
	// 查询的地址
	//baseURL := "http://172.18.6.40:9091"
	//创建的地址
	baseURL := "http://172.18.6.40:9090"
	// 要拼接的值
	value := msg
	// 拼接值到URL
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil, err
	}
	if params == nil {
		fmt.Println("未传递参数！")
	}
	u.Path += "/" + value
	//fmt.Println("url:", u)
	if method == "Get" {
		resp, err := http.Get(u.String())
		if err != nil {
			fmt.Println("发送请求失败:", err)
			return nil, err
		}

		defer resp.Body.Close()
		// 读取响应的 JSON 数据
		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取原始响应数据失败:", err)
			return nil, err
		}
		fmt.Println("原始响应数据:", string(rawResponse))
		return string(rawResponse), nil
	} else if method == "Delete" {
		// 发送DELETE请求
		req, err := http.NewRequest("DELETE", u.String(), nil)
		if err != nil {
			fmt.Println("创建请求失败:", err)
			return nil, err
		}
		// 发送请求
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println("发送请求失败:", err)
			return nil, err
		}
		defer resp.Body.Close()
		// 读取响应的 JSON 数据
		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取原始响应数据失败:", err)
			return nil, err
		}
		fmt.Println("原始响应数据--删除:", string(rawResponse))
		return string(rawResponse), nil

	} else if method == "Post" {
		// 将参数转换为 JSON 格式
		jsonData, err := json.Marshal(params)
		if err != nil {
			fmt.Println("JSON 编码失败:", err)
			return nil, err
		}
		// 发送 POST 请求
		resp, err := http.Post(u.String(), "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("发送请求失败:", err)
			return nil, err
		}
		defer resp.Body.Close()

		// 读取响应的 JSON 数据
		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取原始响应数据失败:", err)
			return nil, err
		}
		fmt.Println("原始响应数据:", string(rawResponse))
		return string(rawResponse), nil
	} else if method == "Put" {
		// 将参数转换为 JSON 格式
		jsonData, err := json.Marshal(params)
		if err != nil {
			fmt.Println("JSON 编码失败:", err)
			return nil, err
		}

		// 创建 PUT 请求
		req, err := http.NewRequest("PUT", u.String(), bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("创建请求失败:", err)
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("发送请求失败:", err)
			return nil, err
		}
		defer resp.Body.Close()

		// 读取响应的 JSON 数据
		rawResponse, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("读取原始响应数据失败:", err)
			return nil, err
		}
		fmt.Println("原始响应数据:", string(rawResponse))
		return string(rawResponse), nil
	} else {
		return nil, nil
	}
}

// 先处理第一个包，主要是返回header，用于后续处理
func ProcFistPackage(data []byte, protocolName string) (msgHeader model.MessageHeader, isPendingRec bool) {
	env := new(model.MessageEnvelope)
	env.ToObject(data)
	header := new(model.MessageHeader)
	header.ToObject(data)
	if protocolName == "UDP" {
		if env.SequenceNumber == 0 && env.MessageLen > 512 {
			return *header, true
		} else {
			return *header, false
		}
	} else {
		if env.SequenceNumber == 0 {
			return *header, true
		} else {
			return *header, false
		}
	}

}

func ResolveBodyData(bodyData []byte, opcode uint32) *model.MessageBody {
	body := new(model.MessageBody)
	body.ToObject(bodyData, opcode)
	return body
}

func unEnvpack(msg []byte) (envObj model.MessageEnvelope) {
	var env model.MessageEnvelope
	env.ToObject(msg)
	return env
}

func unHeaderpack(msg []byte) (envObj model.MessageHeader) {
	var header model.MessageHeader
	header.ToObject(msg)
	return header
}

func unBodypack(msg []byte) (envObj model.UTF8String) {
	var str model.UTF8String
	str.ToObject(msg)
	return str
}
