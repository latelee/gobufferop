package main

import (
	"log"
    "mybuffer"
)

type DataInfo struct {
    head uint8
    len uint16
    len1 uint32
    date string
    name string
    code int
    year int
    
}

/*
十进制读取测试
示例报文（小端）：
4C456400640001002009110768656C6C6F776F726C64992019
解释：
4C 45 : 头部，2个单独字节
6400 长度 16位
64000100 长度 32位
20091107 日期，4字节 BCD码字符串
68656C6C6F776F726C64 字符串，10字节
99 代号 1字节，BCD码，数字
2019： 年份，2字节，BCD码，数字

执行结果：
 head: 0x4c
 len: 100
 len1: 65636 
 date: 20091107 
 name: helloworld 
 code: 99 
 year: 2019

*/
func BufferReaderTest() {
    strData := "4C456400640001002009110768656C6C6F776F726C64992019";

    var reader mybuffer.BufferReader; //{[]byte(mybuffer.ToHexByte(strData)), 0;

    reader.Init([]byte(mybuffer.ToHexByte(strData)));

    var info DataInfo;

    // 读取头
    info.head = reader.ReadUint8();
    // 第二个头跳过
    reader.SkipBytes(1);

    // 长度，分别为2字节、4字节
    info.len = reader.ReadUint16();
    info.len1 = reader.ReadUint32();
    
    // BCD字符串
    info.date = reader.ReadBCDString(4);
    
    // 正常的字符串
    info.name = reader.ReadString(10);
    // 单个bcd码
    info.code = reader.ReadBCD();
    // bcd数值
    info.year = reader.ReadBCDNumber(2);
    log.Printf("data: %#v\n", info); // 直接输出整体结构体
}

func BufferWriterTest() {
    var writer mybuffer.BufferWriter;
    writer.Init(100);
    
    // 按数值写
    writer.WriteUint8(0x4c);
    writer.WriteUint8(0x45);
    writer.WriteUint16(100);
    writer.WriteUint32(0x010064);
    writer.WriteBCDS("20091107");
    
    writer.WriteString("helloworld"); // 正常字符串
    writer.WriteBCD(99);

    writer.WriteBCDS("2019");
    
    p := make([]byte, 2);
    p[0] = 0x7e;
    p[1] = 0x7f;
    writer.WriteBuffer(p); // 原始字符
    
    writer.Test();
}

func main() {                                              
    //log.SetPrefix("Server: ")
    // 设置标志，有log.LstdFlags | log.Lshortfile | 
    log.SetFlags(log.Ldate | log.Lmicroseconds)
    
    //BufferReaderTest();
    BufferWriterTest();
}