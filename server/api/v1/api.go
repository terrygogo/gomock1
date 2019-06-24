package v1

import (
	// "encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	// "time"

	"github.com/mackerelio/go-osstat/network"
	"io/ioutil"
	"net/http"
	//"os"
	//"bytes"
	"github.com/go-chi/chi"
	//"github.com/go-chi/render"
	"log"
	"runtime"
	"strconv"
	"strings"
)

// ValidBearer is a hardcoded bearer token for demonstration purposes.
const ValidBearer = "123456"

// HelloResponse is the JSON representation for a customized message
type HelloResponse struct {
	Message string `json:"message"`
}

// HelloWorld returns a basic "Hello World!" message
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello world!",
	}
	jsonResponse(w, response, http.StatusOK)
}

// HelloName returns a personalized JSON message
func HelloName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	response := HelloResponse{
		Message: fmt.Sprintf("Hello %s!", name),
	}
	jsonResponse(w, response, http.StatusOK)
}

// FruitBasket dff
type FruitBasket struct {
	Name  string
	Fruit []string

	private string // An unexported field is not encoded.

}

// GetJSONFile open any file
func GetJSONFile(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Println(name)
	// Open our jsonFile
	//  dir, err := os.Getwd()
	//  fmt.Println(dir)
	// jsonFile, err := os.Open("./server/data/user.json")
	// if we os.Open returns an error then handle it
	// if err != nil {
	// 	fmt.Println(err)
	//}

	// defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	byteValue, _ := ioutil.ReadFile("./server/data/user.json")
	/*
		basket := FruitBasket{
			Name:  "Standard",
			Fruit: []string{"Apple", "Banana", "Orange"},

			private: "Second-rate",
		}

		var jsonData []byte
		jsonData, err := json.Marshal(basket)
		if err != nil {

		}
		fmt.Println(string(jsonData))*/
	jsonResponse(w, string(byteValue), http.StatusOK)
	/*
			enableCors(&w)
			bp := new(bytes.Buffer)
			if err := json.Compact(bp, byteValue); err != nil {
				fmt.Println(err)
			}
			papaa := fmt.Sprintf("%s", bp)
			mama, err := strconv.Unquote(papaa)
		        if ( err != nil ) {
		           fmt.Println( err) }
			render.JSON(w, r, mama) */
}

type test_struct struct {
	FirstName string
	LastName  string
}

// SetJSONFile ff
func SetJSONFile(w http.ResponseWriter, r *http.Request) {
	//  data := r.Context().Value("")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", reqBody)
	/*
		  	decoder := json.NewDecoder(r.Body)
		    bb := r.Body
			var t test_struct
			err := decoder.Decode(&t)

			if err != nil {
				panic(err)
			}
		  	fmt.Println(bb)
			fmt.Println(t.FirstName)
	*/
	// name := chi.URLParam(r, "name")

	// Open our jsonFile
	//  dir, err := os.Getwd()
	//  fmt.Println(dir)
	// jsonFile, err := os.Open("./server/data/user.json")
	// if we os.Open returns an error then handle it
	// if err != nil {
	// 	fmt.Println(err)
	//}

	// defer the closing of our jsonFile so that we can parse it later on
	// defer jsonFile.Close()

	err1 := ioutil.WriteFile("./server/data/user1.json", []byte(reqBody), 0644)
	if err1 != nil {
		log.Fatal(err1)
	}
	byteValue, _ := ioutil.ReadFile("./server/data/user1.json")
	/*
		basket := FruitBasket{
			Name:  "Standard",
			Fruit: []string{"Apple", "Banana", "Orange"},

			private: "Second-rate",
		}

		var jsonData []byte
		jsonData, err := json.Marshal(basket)
		if err != nil {

		}
		fmt.Println(string(jsonData))*/
	jsonResponse(w, string(byteValue), http.StatusOK)

}

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

// ServerStats define rerun fields
type ServerStats struct {
	HostName         string
	Uptime           string
	NumOfProcess     string
	CPUUsed          string
	MemoryUsed       string
	DiskUsed         string
	NetworkBytesSent string
	NetworkBytesRead string
	RxBytes          string
	TxBytes          string
}

// ServerResourceStats return server resource statat in JSON
func ServerResourceStats(w http.ResponseWriter, r *http.Request) {

	rj := new(ServerStats)

	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	hostStat, err := host.Info()
	dealwithErr(err)
	diskStat, err := disk.Usage("/")

	dealwithErr(err)

	// cpu - get CPU number of cores and speed

	cpercentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	rj.MemoryUsed = strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64)
	// get disk serial number.... strange... not available from disk package at compile time
	// undefined: disk.GetDiskSerialNumber
	//serial := disk.GetDiskSerialNumber("/dev/sda")

	//html = html + "Disk serial number: " + serial + "<br>"

	rj.DiskUsed = strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64)

	rj.CPUUsed = strconv.FormatFloat(cpercentage[0], 'f', 2, 64)

	rj.HostName = hostStat.Hostname
	rj.Uptime = strconv.FormatUint(hostStat.Uptime, 10)
	rj.NumOfProcess = strconv.FormatUint(hostStat.Procs, 10)
	netiocounter, err := net.IOCounters(false)
	rj.NetworkBytesRead = strconv.FormatUint(netiocounter[0].BytesRecv, 10)
	rj.NetworkBytesSent = strconv.FormatUint(netiocounter[0].BytesSent, 10)
	before, err := network.Get()
	if err != nil {

		return
	}
	//time.Sleep(time.Duration(1) * time.Second)
//	after, err := network.Get()
	//if err != nil {

	//	return
//	}
	//rj.RxBytes = strconv.FormatUint((after[3].RxBytes-before[3].RxBytes), 10)
	//rj.TxBytes = strconv.FormatUint((after[3].TxBytes-before[3].TxBytes), 10)
		rj.RxBytes = strconv.FormatUint(( before[3].RxBytes), 10)
	rj.TxBytes = strconv.FormatUint(( before[3].TxBytes), 10)
	jsonResponseStruct(w, rj, http.StatusOK)
}

//GetSystemMeters ss
func GetSystemMeters(w http.ResponseWriter, r *http.Request) {
	runtimeOS := runtime.GOOS
	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)

	// disk - start from "/" mount point for Linux
	// might have to change for Windows!!
	// don't have a Window to test this out, if detect OS == windows
	// then use "\" instead of "/"

	diskStat, err := disk.Usage("/")
	dealwithErr(err)

	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealwithErr(err)

	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	dealwithErr(err)

	html := "<html>OS : " + runtimeOS + "<br>"
	html = html + "Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes <br>"
	html = html + "Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes<br>"
	html = html + "Percentage used memory: " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%<br>"

	// get disk serial number.... strange... not available from disk package at compile time
	// undefined: disk.GetDiskSerialNumber
	//serial := disk.GetDiskSerialNumber("/dev/sda")

	//html = html + "Disk serial number: " + serial + "<br>"

	html = html + "Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes <br>"
	html = html + "Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes<br>"
	html = html + "Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes<br>"
	html = html + "Percentage disk space usage: " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%<br>"

	// since my machine has one CPU, I'll use the 0 index
	// if your machine has more than 1 CPU, use the correct index
	// to get the proper data
	html = html + "CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10) + "<br>"
	html = html + "VendorID: " + cpuStat[0].VendorID + "<br>"
	html = html + "Family: " + cpuStat[0].Family + "<br>"
	html = html + "Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10) + "<br>"
	html = html + "Model Name: " + cpuStat[0].ModelName + "<br>"
	html = html + "Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz <br>"

	for idx, cpupercent := range percentage {
		html = html + "Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%<br>"
	}

	html = html + "Hostname: " + hostStat.Hostname + "<br>"
	html = html + "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "<br>"
	html = html + "Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10) + "<br>"

	// another way to get the operating system name
	// both darwin for Mac OSX, For Linux, can be ubuntu as platform
	// and linux for OS

	html = html + "OS: " + hostStat.OS + "<br>"
	html = html + "Platform: " + hostStat.Platform + "<br>"

	// the unique hardware id for this machine
	html = html + "Host ID(uuid): " + hostStat.HostID + "<br>"

	for _, interf := range interfStat {
		mama := interf.String()
		fmt.Println(mama)
		html = html + "------------------------------------------------------<br>"
		html = html + "Interface Name: " + interf.Name + "<br>"

		if interf.HardwareAddr != "" {
			html = html + "Hardware(MAC) Address: " + interf.HardwareAddr + "<br>"
		}

		for _, flag := range interf.Flags {
			html = html + "Interface behavior or flags: " + flag + "<br>"
		}

		for _, addr := range interf.Addrs {
			html = html + "IPv6 or IPv4 addresses: " + addr.String() + "<br>"

		}

	}

	html = html + "</html>"

	w.Write([]byte(html))

}

// RequireAuthentication is an example middleware handler that checks for a
// hardcoded bearer token. This can be used to verify session cookies, JWTs
// and more.
func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Make sure an Authorization header was provided
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		// This is where token validation would be done. For this boilerplate,
		// we just check and make sure the token matches a hardcoded string
		if token != ValidBearer {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		// Assuming that passed, we can execute the authenticated handler
		next.ServeHTTP(w, r)
	})
}

// NewRouter returns an HTTP handler that implements the routes for the API
func NewRouter() http.Handler {
	r := chi.NewRouter()

	// r.Use(RequireAuthentication)

	// Register the API routes
	r.Get("/", HelloWorld)
	r.Get("/getsystemmeters", GetSystemMeters)
	r.Get("/getjsonfil/{name}", GetJSONFile)
	r.Post("/setjsonfil", SetJSONFile)
	r.Get("/getserverresourcestates", ServerResourceStats)

	r.Get("/{name}", HelloName)

	return r
}
