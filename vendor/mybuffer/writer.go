package mybuffer

/*
默认为小端格式
TODO：有符号数、浮点数
*/
import (
    "fmt"
    //"io"
    //"strconv"
    //"bytes"
    "encoding/hex"
    //"encoding/binary"
    
)

type BufferWriter struct {
    buffer []byte
    offset int
}

func (b *BufferWriter) Buffer() ([]byte) {
    return b.buffer;
}

func (b *BufferWriter) Init(n int) {
    b.buffer = make([]byte, n);
    b.offset = 0;
}

func (b *BufferWriter) Test() {

    fmt.Printf("int BufferWriter \n%v \n%d\n", hex.Dump(b.buffer), b.offset);
}

func (b *BufferWriter) WriteUint8(i uint8) {
    b.buffer[b.offset] = i;
    b.offset += 1;
}

func (b *BufferWriter) WriteUint16(i uint16) {
    b.buffer[b.offset] = byte((i)&0xff);
    b.buffer[b.offset+1] = byte((i>>8)&0xff);
    b.offset += 2;
}

func (b *BufferWriter) WriteUint16BE(i uint16) {
    b.buffer[b.offset] = byte((i>>8)&0xff);
    b.buffer[b.offset+1] = byte((i)&0xff);
    b.offset += 2;
}

func (b *BufferWriter) WriteUint32(i uint32) {
    b.buffer[b.offset] = byte((i)&0xff);
    b.buffer[b.offset+1] = byte((i>>8)&0xff);
    b.buffer[b.offset+2] = byte((i>>16)&0xff);
    b.buffer[b.offset+3] = byte((i>>24)&0xff);
    b.offset += 4;
}

func (b *BufferWriter) WriteUint32BE(i uint32) {
    b.buffer[b.offset] = byte((i>>24)&0xff);
    b.buffer[b.offset+1] = byte((i>>16)&0xff);
    b.buffer[b.offset+2] = byte((i>>8)&0xff);
    b.buffer[b.offset+3] = byte((i)&0xff);
    b.offset += 4;
}

func (b *BufferWriter) WriteString(i string) {
    copy(b.buffer[b.offset:], i); // 直接用copy赋值
    b.offset += len(i);
}

func (b *BufferWriter) WriteHexString(i string) {
    ii := ToHexByte(i) // 先转hex字符
    copy(b.buffer[b.offset:], ii); // 直接用copy赋值
    b.offset += len(ii);
}

func (b *BufferWriter) WriteBuffer(i []byte) {
    copy(b.buffer[b.offset:], i); // 直接用copy赋值
    b.offset += len(i);
}

// 一个BCD码最大为99
func (b *BufferWriter) WriteBCD(i uint8) {
    c1 := (i/10)*16 + (i%10);
    //fmt.Printf("%d\n", c1);
    b.buffer[b.offset] = byte(c1);
    b.offset += 1;
}

// 直接以十六进制形式写即可
func (b *BufferWriter) WriteBCDS(bcds string) {
    b.WriteHexString(bcds);
}

/*
// 转入十六进制形式的字符，输出为byte类型
func ToHexByte(str string) (ob []byte) {
    ob, _ = hex.DecodeString(str);
    
    return;
}

// 转入十六进制形式的数组，输出为对应的字符
// 如 4c 77数组，将转换成4c77字符串，可保存到文件
func ToHexString(b []byte) (ostr string) {
    ostr = hex.EncodeToString(b);
    return;
}
*/

