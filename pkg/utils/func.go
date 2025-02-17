//Created by dolitTeam
//@Author : GJing
//@Time : 2020/10/23 13:56
//@File : functions
//@Description: 公共函数库

package utils

import (
	"archive/zip"
	"bufio"
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gjing1st/hertz-admin/internal/pkg/functions"
	"github.com/gjing1st/hertz-admin/pkg/utils/global"
	log "github.com/sirupsen/logrus"
	"io"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

// Md5
// Author: GJing
// Email: gjing1st@gmail.com
// Date: 2020/10/23 13:57
// Description: md5加密
func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

// InArray
//
//	@description:	判断一个字符串是否在数组中
//	@param:
//	@author:	GJing
//	@email:		gjing1st@gmail.com
//	@date:		2020/11/13 14:42
//	@success:
func InArray(value string, arr []string) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

// ReserveNumber
//
//	@description:	截取保留小数点后m位，舍去后面位数
//	@param:			f float64
//	@param:			m 保留的位数
//	@author:		GJing
//	@email:			gjing1st@gmail.com
//	@date:			2020/11/25 上午 10:31
//	@success:		返回截取后的字符串
func ReserveNumber(f float64, m int) string {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[:i+m+1]
		}
	}
	return s
}

// @description:	压缩文件夹
// @param:			dir 文件夹路径 ex:F:\project\Go\ChineseMedicine\ChineseMedicine\adminApi\public\image
// @param:			zipFile 压缩后的文件夹路径和名称 ex: ./test.zip
// @author:		GJing
// @email:			gjing1st@gmail.com
// @date:			2021/1/13 15:17
// @success:
// @remark:	相对路径压缩后可能导致里面目录名称错误，可使用绝对路径。具体原因未知。str, _ := os.Getwd()获取当前程序运行所在目录，str拼接相对路径
func zipDir(dir, zipFile string) {
	// TODO 此加解压有问题，使用中医中最新的加解压
	fz, err := os.Create(zipFile)
	if err != nil {
		log.Fatalf("Create zip file failed: %s\n", err.Error())
	}
	defer fz.Close()

	w := zip.NewWriter(fz)
	defer w.Close()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fDest, err := w.Create(path[len(dir)+1:])
			if err != nil {
				log.Printf("Create failed: %s\n", err.Error())
				return nil
			}
			fSrc, err := os.Open(path)
			if err != nil {
				log.Printf("Open failed: %s\n", err.Error())
				return nil
			}
			defer fSrc.Close()
			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				log.Printf("Copy failed: %s\n", err.Error())
				return nil
			}
		}
		return nil
	})
}

// UnzipDir
//
//	@description:	解压缩
//	@param:zipFile	压缩文件路径 ./test.zip
//	@param:dir		需要解压到的指定文件夹目录 ex :F:\dumps_copy
//	@author:		GJing
//	@email:			gjing1st@gmail.com
//	@date:			2021/1/13 15:19
//	@success:
func UnzipDir(zipFile, dir string) {

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Println("zipFile", zipFile)
		log.Fatalf("Open zip file failed: %s\n", err.Error())
	}
	defer r.Close()

	for _, f := range r.File {
		func() {
			path := dir + string(filepath.Separator) + f.Name
			os.MkdirAll(filepath.Dir(path), 0755)
			fDest, err := os.Create(path)
			if err != nil {
				log.Printf("Create failed: %s\n", err.Error())
				return
			}
			defer fDest.Close()

			fSrc, err := f.Open()
			if err != nil {
				log.Printf("Open failed: %s\n", err.Error())
				return
			}
			defer fSrc.Close()

			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				log.Printf("Copy failed: %s\n", err.Error())
				return
			}
		}()
	}
}

// Round
//
//	@description:	四舍五入保留n位小数
//	@param:f		需要处理的float数
//	@param:n		需要保留的小数位数
//	@author:		GJing
//	@email:			gjing1st@gmail.com
//	@date:			2021/1/23 10:36
//	@success:
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

// Div 数字转字母
func Div(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}
	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}

// UnExt
//
//	@description:	返回文件名称去掉后缀和最后一个`.`
//	@param:			fileName 文件名称
//	@author:		GJing
//	@email:			gjing1st@gmail.com
//	@date:			2022/7/28 15:39
//	@success:
func UnExt(fileName string) string {
	for i := len(fileName) - 1; i >= 0 && fileName[i] != '/'; i-- {
		if fileName[i] == '.' {
			return fileName[:i]
		}
	}
	return ""
}

type connection struct {
	mu        sync.Mutex
	sshclient *ssh.Client
}

func (c *connection) session() (*ssh.Session, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.sshclient == nil {
		return nil, errors.New("connection closed")
	}

	sess, err := c.sshclient.NewSession()
	if err != nil {
		return nil, err
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = sess.RequestPty("xterm", 100, 50, modes)
	if err != nil {
		return nil, err
	}

	return sess, nil
}

type Host struct {
	IP       string
	Password string
	User     string
}

func (c *connection) Exec(cmd string, host Host) (stdout string, code int, err error) {
	sess, err := c.session()
	if err != nil {
		return "", 1, errors.New("failed to get SSH session")
	}
	defer sess.Close()

	exitCode := 0

	in, _ := sess.StdinPipe()
	out, _ := sess.StdoutPipe()

	err = sess.Start(strings.TrimSpace(cmd))
	if err != nil {
		exitCode = -1
		if exitErr, ok := err.(*ssh.ExitError); ok {
			exitCode = exitErr.ExitStatus()
		}
		return "", exitCode, err
	}

	var (
		output []byte
		line   = ""
		r      = bufio.NewReader(out)
	)

	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}

		output = append(output, b)

		if b == byte('\n') {
			line = ""
			continue
		}

		line += string(b)

		if (strings.HasPrefix(line, "[sudo] password for ") || strings.HasPrefix(line, "Password")) && strings.HasSuffix(line, ": ") {
			_, err = in.Write([]byte(host.Password + "\n"))
			if err != nil {
				break
			}
		}
	}
	err = sess.Wait()
	if err != nil {
		exitCode = -1
		if exitErr, ok := err.(*ssh.ExitError); ok {
			exitCode = exitErr.ExitStatus()
		}
	}
	outStr := strings.TrimPrefix(string(output), fmt.Sprintf("[sudo] password for %s:", host.User))

	// preserve original error
	return strings.TrimSpace(outStr), exitCode, errors.New(fmt.Sprintf("Failed to exec command: %s \n%s", cmd, strings.TrimSpace(outStr)))
}

// RunCommand
//
//	@description:	运行系统命令
//	@param:			cmdStr 要运行的命令
//	@author:		GJing
//	@email:			gjing1st@gmail.com
//	@date:			2022/9/2 17:46
//	@success:
func RunCommand(name string, arg ...string) (err error) {
	cmd := exec.Command(name, arg...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		log.Println("运行系统命令错误", err, ":", stderr.String())
		return
	}
	return
}

// WriteFile
//
//	@description:
//	@param:
//	@author:	GJing
//	@email:		gjing1st@gmail.com
//	@date:		2022/9/2 17:52
//	@success:
func WriteFile(fileName, s string) (err error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Println("文件打开失败", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(s)
	write.Flush()
	return
}

// PathExists
//
//	@description:	判断文件是否存在
//	@param:
//	@author:	Zq
//	@email:		zhengqiang@tna.cn
//	@date:		2022/10/19 17:52
//	@success:
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// LogField
//
//	@description:	日志字段
//	@param:
//	@author:	GJing
//	@email:		gjing1st@gmail.com
//	@date:		2022/11/1 15:35
//	@success:
func LogField(err error, msg string) log.Fields {
	return log.Fields{
		"err": err,
		"msg": msg,
	}
}

// DiffNatureDays
//
//	@description:	两个日期的相差的自然天数
//	@param:
//	@author:	GJing
//	@email:		gjing1st@gmail.com
//	@date:		2022/11/16 11:22
//	@success:
func DiffNatureDays(t1, t2 int64) int {
	if t1 == t2 {
		return -1
	}
	if t1 > t2 {
		t1, t2 = t2, t1
	}

	diffDays := 0
	secDiff := t2 - t1
	if secDiff > global.SecondsPerDay {
		tmpDays := int(secDiff / global.SecondsPerDay)
		t1 += int64(tmpDays) * global.SecondsPerDay
		diffDays += tmpDays
	}

	st := time.Unix(t1, 0)
	et := time.Unix(t2, 0)
	dateFormatTpl := "20060102"
	if st.Format(dateFormatTpl) != et.Format(dateFormatTpl) {
		diffDays += 1
	}

	return diffDays
}

// DockerRunCommand
// @description: docker容器执行宿主机指令
// @param:
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/11 13:46
// @success:
func DockerRunCommand(command string) (err error) {
	/*
	   nsenter命令是一个可以在指定进程的命令空间下运行指定程序的命令，位于util-linux包中  。该命令作为容器向宿主机发送命令的关键部分 。
	   使用格式：nsenter -a -t <pid> <command> 或者nsenter -m -u -i -n -p -t <pid> <command> ；
	    -a表示进入宿主机的所有命名空间 , -t 表示获取/proc/{pid}进程 ，liniux旧版本可能不支持。
	   需要使用 -m -u -i -n -p 。 -m -u -i -n -p，表示进入mount, UTS,System V IPC,网络,pid命名空间,
	   这几个命名空间包含了绝大多数的空间环境。
	*/
	cmd := exec.Command("nsenter", "-m", "-u", "-i", "-n", "-p", "-t", "1", "sh", "-c", command)
	log.Debug("DockerRunCommand==", cmd.String())
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		//log.Println("运行系统命令错误", err, ":", stderr.String())
		functions.AddErrLog(log.Fields{"msg": "运行系统命令错误", "cmd": cmd.String(), "err": stderr.String()})
		return
	}
	return
}

// VersionCompare
// @description: 版本比较，v2版本是否大于v1版本
// @param: v1 string 当前版本
// @param: v2 string 最新版本
// @author: GJing
// @email: gjing1st@gmail.com
// @date: 2023/2/16 14:41
// @success: v2>v1返回true
func VersionCompare(v1, v2 string) (res bool) {
	if len(v1) == 0 {
		return true
	}
	if len(v2) == 0 {
		return false
	}
	if v1[0] == 'v' || v1[0] == 'V' {
		v1 = v1[1:]
	}
	if v2[0] == 'v' || v2[0] == 'V' {
		v2 = v2[1:]
	}
	v1Arr := strings.Split(v1, ".")
	v2Arr := strings.Split(v2, ".")
	//判断是否可升级
	if len(v1Arr) >= 3 && len(v2Arr) >= 3 {
		if Int(v2Arr[0]) > Int(v1Arr[0]) {
			res = true
			return
		} else if Int(v2Arr[0]) < Int(v1Arr[0]) {
			return
		}
		if Int(v2Arr[1]) > Int(v1Arr[1]) {
			res = true
			return
		} else if Int(v2Arr[1]) < Int(v1Arr[1]) {
			return
		}
		if Int(v2Arr[2]) > Int(v1Arr[2]) {
			res = true
			return
		} else if Int(v2Arr[2]) < Int(v1Arr[2]) {
			return
		}
		if Int(v2Arr[3]) > Int(v1Arr[3]) {
			res = true
			return
		}
	}
	return
}

// IsValidIP
// @Description 验证IP是否有效
// @params
// @contact.name GJing
// @contact.email gjing1st@gmail.com
// @date 2023/6/12 18:41
func IsValidIP(ip string) bool {
	pattern := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)
	return pattern.MatchString(ip)
}
