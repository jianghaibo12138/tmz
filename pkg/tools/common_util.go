package tools

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const FormatStr = "2006-01-02 15:04:05"
const FormatTimeSuffix = "20060102150405"
const FormatDate = "2006-01-02"
const FormatTime = "15:04:05"
const FormatStrExcel = "2006/01/02 15:04:05.0"

//
// ContainsString
// @Description: 查询指定字符串在数组中的位置
// @param array
// @param val
// @return index
//
func ContainsString(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return index
		}
	}
	return index
}

//
// IsContainsString
// @Description: 通过正则查询字符串是否包含在数组中，包含返回对应的索引，不包含返回-1
// @param array
// @param val
// @return index
//
func IsContainsString(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		var pattern = fmt.Sprintf("%s", array[i])
		res, _ := regexp.MatchString(pattern, val)
		if res {
			index = i
			return index
		}
	}
	return index
}

//
// ContainsInt
// @Description: 查询指定数字在数组中的位置（int）
// @param array
// @param val
// @return index
//
func ContainsInt(array []int, val int) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return index
		}
	}
	return index
}

//
// ContainsInt64
// @Description: 查询指定数字在数组中的位置（int64）
// @param array
// @param val
// @return index
//
func ContainsInt64(array []int64, val int64) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return index
		}
	}
	return index
}

//
// ContainsUint
// @Description: 查询指定数字在数组中的位置（uint64）
// @param array
// @param val
// @return index
//
func ContainsUint(array []uint64, val uint64) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return index
		}
	}
	return index
}

//
// ContainsFloat
// @Description: 查询指定数字在数组中的位置（float64）
// @param array
// @param val
// @return index
//
func ContainsFloat(array []float64, val float64) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return index
		}
	}
	return index
}

// ContainsComplex Returns the index position of the complex128 val in array
//
// ContainsComplex
// @Description: 查询指定数字在数组中的位置（complex128）
// @param array
// @param val
// @return index
//
func ContainsComplex(array []complex128, val complex128) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return index
		}
	}
	return index
}

//
// Exists
// @Description: 判断所给路径文件/文件夹是否存在
// @param path
// @return bool
//
func Exists(path string) bool {
	_, err := os.Stat(path) // os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//
// IsDir
// @Description: 判断所给路径是否为文件夹
// @param path
// @return bool
//
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//
// IsFile
// @Description: 判断所给路径是否为文件
// @param path
// @return bool
//
func IsFile(path string) bool {
	return !IsDir(path)
}

//
// GetCurDirList
// @Description: 获取指定目录下的所有目录，不进入下一级目录搜索
// @param dirPth
// @return []string
// @return error
//
func GetCurDirList(dirPth string) ([]string, error) {
	var files []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if fi.IsDir() {
			files = append(files, fi.Name())
		}
	}
	return files, nil
}

//
// GetDirList
// @Description: 获取目录下所有的文件夹，包括层级目录下
// @param dirPath
// @return []string
// @return error
//
func GetDirList(dirPath string) ([]string, error) {
	var dirList []string
	err := filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			dirList = append(dirList, path)
			return nil
		}
		return nil
	})
	return dirList, err
}

//
// GetAllFile
// @Description: 获取指定目录下所有的文件名，包括层级目录下
// @param pathName
// @param s
// @return []string
// @return error
//
func GetAllFile(pathName string, s []string) ([]string, error) {
	rd, err := ioutil.ReadDir(pathName)
	if err != nil {
		return s, err
	}
	for _, fi := range rd {
		if fi.IsDir() {
			fullDir := path.Join(pathName, fi.Name())
			s, err = GetAllFile(fullDir, s)
			if err != nil {
				return s, err
			}
		} else {
			// 添加文件限制，防止一下子获取全部文件
			if len(s) > 10000 {
				return s, err
			}
			fullName := path.Join(pathName, fi.Name())
			s = append(s, fullName)
		}
	}
	return s, nil
}

//
// IntZFill
// @Description: 补零函数
// @param source
// @param maxLength
// @return string
//
func IntZFill(source int64, maxLength int) string {
	formatter := "%0" + strconv.Itoa(maxLength) + "d"
	return fmt.Sprintf(formatter, source)
}

//
// IsUpper
// @Description: 判断字符串是否都是大写
// @param s
// @return bool
//
func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

//
// IsLower
// @Description: 判断字符串是否都是小写
// @param s
// @return bool
//
func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

//
// Md5Hash
// @Description: 生成MD5
// @param data
// @return string
//
func Md5Hash(data []byte) string {
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)
	return md5str
}

//
// TimeStr2LocalTime
// @Description: 时间字符串转datetime
// @param string
// @param val
// @return time.Time
//
func TimeStr2LocalTime(timeStr string) (time.Time, error) {
	loc, _ := time.LoadLocation("Local")
	startTime, err := time.ParseInLocation(FormatStr, timeStr, loc)
	return startTime, err
}

//
// TimeStr2Unix
// @Description: 时间字符串转unix
// @param string
// @param val
// @return int64
//
func TimeStr2Unix(timeStr string) int64 {
	var t int64
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation(FormatStr, timeStr, loc)
	t = t1.Unix()
	// fmt.Printf("%v-->:%v \n",time_str, t)
	return t
}

//
// TimeStamp2TimeStr
// @Description: 时间戳数据转时间字符串
// @param t
// @return string
//
func TimeStamp2TimeStr(t int64) string {
	s := time.Unix(t, 0).Format(FormatStrExcel)
	return s
}

//
// ExecShell
// @Description: 执行shell指令函数，需要注意可执行文件是否运行在宿主机器上
// @param cmdStr
// @param args
// @return string
// @return error
//
func ExecShell(cmdStr string, args ...string) (string, error) {
	// 函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(cmdStr, args...)

	// 读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	// Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	return out.String(), err
}

//
// GetDaysOfMonth
// @Description: 获取当前月有多少天
// @param t
// @return []string
//
func GetDaysOfMonth(t time.Time) []string {
	days := make([]string, 0)
	for month := 1; month <= 12; month++ {
		if month != int(t.Month()) {
			continue
		}
		for day := 1; day <= 31; day++ {
			// 如果是2月
			if month == 2 {
				if IsLeapYear(t.Year()) && day == 30 { // 闰年2月29天
					break
				} else if !IsLeapYear(t.Year()) && day == 29 { // 平年2月28天
					break
				} else {
					days = append(days, fmt.Sprintf("%d-%02d-%02d", t.Year(), month, day))
				}
			} else if month == 4 || month == 6 || month == 9 || month == 11 { // 小月踢出来
				if day == 31 {
					break
				}
				days = append(days, fmt.Sprintf("%d-%02d-%02d", t.Year(), month, day))
			} else {
				days = append(days, fmt.Sprintf("%d-%02d-%02d", t.Year(), month, day))
			}
		}
	}
	return days
}

//
// IsLeapYear
// @Description: 是否是闰年
// @param year
// @return bool
// @return 2004
//
func IsLeapYear(year int) bool { // y == 2000, 2004
	// 判断是否为闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}

	return false
}

//
// Str2DateTime
// @Description: 字符串转日期时间函数，例如："2021-01-01 23:59:59"
// @param timeStr
// @return time.Time
// @return error
//
func Str2DateTime(timeStr string) (time.Time, error) {
	var loc, _ = time.LoadLocation("Local")
	dateTime, err := time.ParseInLocation(FormatStr, timeStr, loc)
	return dateTime, err
}

//
// Str2Date
// @Description: 字符串转日期函数，例如： "2021-01-01"
// @param timeStr
// @return time.Time
// @return error
//
func Str2Date(timeStr string) (time.Time, error) {
	var loc, _ = time.LoadLocation("Local")
	dateTime, err := time.ParseInLocation(FormatDate, timeStr, loc)
	return dateTime, err
}

//
// Str2Time
// @Description: 字符串转时间函数，例如："23:59:59"
// @param timeStr
// @return time.Time
// @return error
//
func Str2Time(timeStr string) (time.Time, error) {
	var loc, _ = time.LoadLocation("Local")
	dateTime, err := time.ParseInLocation(FormatTime, timeStr, loc)
	return dateTime, err
}

//
// GetFirstDateOfMonth
// @Description: 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
// @param d
// @return time.Time
//
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetStartTimeOfDay(d)
}

//
// GetStartTimeOfDay
// @Description: 获取某一天的0点时间
// @param d
// @return time.Time
//
func GetStartTimeOfDay(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

//
// GetLastDateOfMonth
// @Description:  获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
// @param d
// @return time.Time
//
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

//
// compress
// @Description: 压缩执行函数
// @param file
// @param prefix
// @param zw
// @return error
//
func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		if len(prefix) == 0 {
			prefix = info.Name()
		} else {
			prefix = prefix + "/" + info.Name()
		}
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		if len(prefix) == 0 {
			header.Name = header.Name
		} else {
			header.Name = prefix + "/" + header.Name
		}
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//
// Zip
// @Description: 压缩函数，源文件可以是一个文件或者目录
// @param srcFile
// @param destZip
// @return error
//
func Zip(srcFile string, destZip string) error {
	var spec string
	if runtime.GOOS == "windows" {
		spec = "\\"
	} else {
		spec = "/"
	}
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	filepath.Walk(srcFile, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(srcFile)+spec)
		// header.Name = path
		if info.IsDir() {
			header.Name += spec
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	return err
}

//
// ZipDir 打包成zip文件
// @Description:
// @param srcDir
// @param zipFileName
//
func ZipDir(srcDir string, zipFileName string) {

	// 预防：旧文件无法覆盖
	os.RemoveAll(zipFileName)

	// 创建：zip文件
	zipFile, _ := os.Create(zipFileName)
	defer zipFile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	var dirSpec string
	if runtime.GOOS == "windows" {
		dirSpec = "\\"
	} else {
		dirSpec = "/"
	}

	// 遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+dirSpec)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += dirSpec
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})
}

//
// GetIpv4Address
// @Description: 获IPv4地址
// @return string
// @return error
//
func GetIpv4Address() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String(), nil
			}
		}
	}
	return "", nil
}

//
// Ipv42Int
// @Description: 将IPv4地址转换为int64
// @param ip
// @return int64
//
func Ipv42Int(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}
