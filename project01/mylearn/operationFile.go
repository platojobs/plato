package mylearn

import (
	"bytes"
	"fmt"
	"io"

	//"log"
	"bufio"
	"encoding/xml"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"os"
	"net"
	"strings"
)

// 操作文件-写入内容
func OperationFile() {
	file, eror := os.OpenFile("mylearn/poetry.txt", os.O_RDWR|os.O_APPEND, 0666)
	if eror != nil {
		fmt.Println("打开文件失败", eror)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("锄禾日当午\n")
	write.WriteString("汗滴禾下土\n")
	write.WriteString("谁知盘中餐\n")
	write.WriteString("粒粒皆辛苦\n")
	write.Flush()
	fmt.Println("写入成功")
}

// 追加内容
func AppendFile() {
	file, eror := os.OpenFile("mylearn/poetry.txt", os.O_RDWR|os.O_APPEND, 0666)
	if eror != nil {
		fmt.Println("打开文件失败", eror)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString("  \n")
	write.WriteString("昨夜雨疏风骤浓睡不消残酒\n")
	write.WriteString("试问卷帘人却道海棠依旧知否知否\n")
	write.Flush()
	fmt.Println("追加诗句成功写入成功")
}

func ReadFile() {
	file, eror := os.Open("mylearn/poetry.txt")
	if eror != nil {
		fmt.Println("打开文件失败", eror)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(readString)
	}
}

func GobEncode() {
	info := `Go语言中的encoding/gob包也提供了与encoding/json包一样的编码、解码功能。
	 通常而言,如果对文件的可读性没有要求,gob格式是Go语言中文件存储和网络传输最方便的格式.`
	file, orr := os.Create("mylearn/hello_go.gob")
	if orr != nil {
		fmt.Println("创建文件失败", orr)
		return
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err := encoder.Encode(info)
	if err != nil {
		fmt.Println("编码失败", err)
		return
	}
	fmt.Println("编码成功")

}

func GobDecode() {
	file, err := os.Open("mylearn/hello_go.gob")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	var info string
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err)
		return
	}
	fmt.Println("解码成功")
	fmt.Println(info)
}

//自定义二进制格式

type WebSite struct {
	Uri int64
}

func BinaryEncode() {
	file, err := os.Create("mylearn/hello_go.bin")
	for i := 0; i < 10; i++ {
		info := WebSite{int64(i)}
		if err != nil {
			fmt.Println("文件创建失败", err)
			return
		}
		defer file.Close()
		var bin_buf bytes.Buffer
		binary.Write(&bin_buf, binary.LittleEndian, info)
		b := bin_buf.Bytes()
		_, beer := file.Write(b)
		if beer != nil {
			fmt.Println("编码失败", beer.Error())
			return
		}
	}
	fmt.Println("编码成功")

}

func BinaryDecode() {

	file, err := os.Open("mylearn/hello_go.bin")
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	info := WebSite{}
	for i := 0; i < 10; i++ {
		data := readNextBytes(file, 8)
		buf := bytes.NewBuffer(data)
		err = binary.Read(buf, binary.LittleEndian, &info)
		if err != nil {
			fmt.Println("二进制读取失败", err)
			return
		}
		fmt.Println("第", i, "个值是", info)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}

type Student struct {
	Name string
	Age  int
	Sex  string
}

func JsonEcodeTT() {
	student := []Student{
		{"张三", 18, "男"},
		{"李四", 19, "女"},
		{"王五", 20, "男"},
	}
	file, err := os.Create("mylearn/student.json")

	if err != nil {
		fmt.Println("创建学生文件失败", err)
		return
	}
	defer file.Close()
	//创建编码器
	encoder := json.NewEncoder(file)
	err = encoder.Encode(student)
	if err != nil {
		fmt.Println("学生编码失败", err)
		return
	}
	fmt.Println("学生编码成功")

}

func JsonDecode() {
	file, err := os.Open("mylearn/student.json")
	if err != nil {
		fmt.Println("打开学生文件失败", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var student []Student
	err = decoder.Decode(&student)
	if err != nil {
		fmt.Println("学生解码失败", err)
		return
	}
	fmt.Println("学生解码成功")
	fmt.Println(student)
}

type XmlSite struct {
	Name    string `xml:"name,attr"`
	Url     string
	Content []string
}

func XmlEcode() {

	siteInfo := XmlSite{
		"Go语言中文网", "http://www.topgoer.com", []string{"Go语言", "Golang", "GOPHER"}}
	file, err := os.Create("mylearn/hello_go.xml")
	if err != nil {
		fmt.Println("创建xml文件失败", err)
		return
	}
	defer file.Close()
	encoder := xml.NewEncoder(file)
	err = encoder.Encode(siteInfo)
	if err != nil {
		fmt.Println("xml编码失败", err)
		return
	}
	fmt.Println("xml编码成功")

}

func XmlDecode(){

	file, err := os.Open("mylearn/hello_go.xml")
	if err != nil {
		fmt.Println("打开xml文件失败", err)
		return
	}
	defer file.Close()
	decoder := xml.NewDecoder(file)
	var siteInfo XmlSite
	err = decoder.Decode(&siteInfo)
	if err != nil {
		fmt.Println("xml解码失败", err)
		return
	}
	fmt.Println("xml解码成功")
	fmt.Println(siteInfo)

}


//TCP服务端
func TcpServise(){
	listen,err := net.Listen("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("监听失败",err)
		return
	}
	fmt.Println("监听成功,等待客户端连接请求")
	for {
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("连接失败",err)
			return
		}
		fmt.Println("连接成功")
		go SRprocess(conn)
	}
	
}

func SRprocess(conn net.Conn){
	defer conn.Close()
	
	for {
		var buf [1024]byte
		n,err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("接收失败",err)
			break
		}
		fmt.Println("已成功接收客户端发来消息:",string(buf[:n]))
		if _,cerr := conn.Write([]byte("服务端消息"));cerr != nil {
			fmt.Println("发送失败",cerr)
			break
		}
	}
}

//TCP客户端
func TcpClient(){
	net_conn,err := net.Dial("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接失败",err)
		return
	}
	fmt.Println("客户端开始发送请求")
	defer net_conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for{
      input,err := inputReader.ReadString('\n')
	  if err != nil {
		  fmt.Println("无法读取失败",err)
		  break
	  }
	  trimmmedInput := strings.TrimSpace(input)
      if trimmmedInput == "Q" {
		break
	  }
	  if _,err := net_conn.Write([]byte(trimmmedInput));err != nil {
		  fmt.Println("发送数据失败",err)
		  break
	  }
      recdata := make([]byte,1024)
	
	  if  _,err := net_conn.Read(recdata);err != nil {
		  fmt.Println("接收失败",err)
		  break
	  }
	  fmt.Println("接收数据成功:",string(recdata))
	}
	
}




func UdpServise(){

	udp_addr,err := net.ResolveUDPAddr("udp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("解析地址失败",err)
		return
	}
	udp_conn,err := net.ListenUDP("udp",udp_addr)
	if err != nil {
		fmt.Println("监听失败",err)
		return
	}
	defer udp_conn.Close()
	fmt.Println("监听成功,等待客户端连接请求")
	for {
		var buf [1024]byte
		n,udp_addr,err := udp_conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接收失败",err)
			break
		}
		fmt.Println("已成功接收客户端发来消息:",string(buf[:n]))
		if _,cerr := udp_conn.WriteToUDP([]byte("hello,客服端"),udp_addr);cerr != nil {
			fmt.Println("发送失败",cerr)
			continue
		}
	}
}

func UdpClient(){
	udp_addr,err := net.ResolveUDPAddr("udp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("解析地址失败",err)
		return
	}
	udp_conn,err := net.DialUDP("udp",nil,udp_addr)
	if err != nil {
		fmt.Println("连接失败",err)
		return
	}
	defer udp_conn.Close()
	senddata := []byte("hello,服务端")
	_,scerrs := udp_conn.Write(senddata)
	
	  if scerrs != nil {
		  fmt.Println("无法读取失败",err)
		  return
	  }
	  data := make([]byte,4096)
	  _,_,cerr := udp_conn.ReadFromUDP(data)
	  if cerr != nil {
		  fmt.Println("接收失败",err)
		  return
	  }
	  fmt.Printf("接收数据成功:%s\n",string(data))
}