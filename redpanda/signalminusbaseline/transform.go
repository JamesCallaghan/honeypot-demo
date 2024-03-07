package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/redpanda-data/redpanda/src/transform-sdk/go/transform"
)

func main() {
	// Register your transform function.
	// This is a good place to perform other setup too.
	keysSet = make(map[string]struct{}, len(keys))
	for _, key := range keys {
		keysSet[key] = struct{}{}
	}
	transform.OnRecordWritten(doTransform)
}

type Message struct {
	Timestamp string                 `json:"timestamp"`
	Data      map[string]interface{} `json:",inline"`
}

// To create the "hash" of a known message, focused on avoiding false negatives

var topLevelFields = []string{"process_exec", "process_exit", "process_kprobe"}
var subFieldsToConcatenate = []string{"process.pod.container.id", "process.binary", "process.arguments"}

func createKey(incomingMessage map[string]interface{}) string {
	var keyParts []string

	for _, topLevelField := range topLevelFields {
		if value, ok := incomingMessage[topLevelField]; ok {
			// If the top-level field exists, traverse its subfields
			for _, subField := range subFieldsToConcatenate {
				// Split the subfield into parts
				parts := strings.Split(subField, ".")
				subValue := value.(map[string]interface{})
				for _, part := range parts {
					// Traverse the map
					if v, ok := subValue[part]; ok {
						// If the part exists, add it to the key parts
						switch v := v.(type) {
						case map[string]interface{}:
							subValue = v
						default:
							keyParts = append(keyParts, fmt.Sprint(v))
						}
					}
				}
			}
		}
	}

	// Join the key parts with a separator
	key := strings.Join(keyParts, "")
	// Remove all whitespaces and escape characters
	key = strings.ReplaceAll(key, " ", "")
	key = strings.ReplaceAll(key, "\\", "")
	key = strings.ReplaceAll(key, "/", "")
	key = strings.ReplaceAll(key, "\"", "")
	key = strings.ReplaceAll(key, "'", "")
	key = strings.ReplaceAll(key, "-", "")
	key = strings.ReplaceAll(key, "/", "")
	key = strings.ReplaceAll(key, "=", "")
	key = strings.ReplaceAll(key, ".", "")
	key = strings.ReplaceAll(key, "containerd", "")
	key = strings.ReplaceAll(key, ":", "")
	key = strings.ReplaceAll(key, "+", "")
	key = strings.ReplaceAll(key, "$", "")
	key = strings.ReplaceAll(key, "_", "")

	return key
}

var keys = []string{
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat3337727510version",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat3337727510mmachI3IEbslMT6s",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat3458575442version",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat3458575442mmachIvSMlJqeQOE",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat1415603660version",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat1415603660mmachIvSMlJqeQOE",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinlsofnPEFfgi4i6",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat3632165990version",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat3632165990mmachIvSMlJqeQOE",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat775682583version",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat775682583mmachI3IEbslMT6s",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat3136117077version",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat3136117077mmachI3IEbslMT6s",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinlsofnPEFfgi4i6",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat3968624077version",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat3968624077mmachIvSMlJqeQOE",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat2243628187version",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat2243628187mmachI3IEbslMT6s",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat3616671096version",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eoptspyderbattmpbat3616671096mmachIvSMlJqeQOEp128130202100",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbindpkgqueryf{Package}{Version}",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinunamesrio",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155optspyderbattmpbat12970479version",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinfindhostruninamecriosock",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinfindhostruninamedockershimsock",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinfindhostruninamedockershimsock",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161ebinsh",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbincurlretry5httpsorcspyderbatcomv1regbL6Ns0xJY0MDliFn0xqXscript",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinwhoami",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinuname",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbingrepLinux",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbingrepCentOSrelease6etcsystemrelease",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinunamer",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinawkF{print12}",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinunamem",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinunamea",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinloggerNanoAgentinstallplatformdetectedUbuntudebian220460x8664",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinmkdirpoptspyderbatetc",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinmkdirpoptspyderbattmp",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinmkdirpoptspyderbatbin",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinchmod700optspyderbatetc",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinchmod600optspyderbattmp",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinchmod600optspyderbatbin",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbincat",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinchmod600optspyderbatetcregistration",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinchmod700optspyderbatbinnanoloopsh",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinchmod700optspyderbatbincentos6spyderbat",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbincat",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod600optspyderbatetcregistration",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod700optspyderbatbinnanoloopsh",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod700optspyderbatbincentos6spyderbat",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod600optspyderbatetcconfiguration",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod700optspyderbatbinremovespyderbatsh",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbincurlhelp",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod700optspyderbatbinreinstallsh",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod700optspyderbatbincentos7selinuxsh",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbincurlHAcceptapplicationoctetstreamooptspyderbatbinnahttpsorcspyderbatcomv1regbL6Ns0xJY0MDliFn0xqXagentx8664latest",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinfindhostruninamesock",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod700optspyderbatbinna",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13aoptspyderbatbinnalink",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinmvoptspyderbatbinnaoptspyderbatbinnanoagentv1260",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinlnfsoptspyderbatbinnanoagentv1260optspyderbatbinnanoagent",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13aoptspyderbatbinnanoagentstdoutvurlhttpsorcspyderbatcomagentRegistrationCodebL6Ns0xJY0MDliFn0xqXk8smonitor",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinfindhostruninamesock",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinmkdirprun",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinlnshostrunk3ssockrunsock",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinfindhostruninamecriosock",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13aoptspyderbattmpbat1648488918version",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13aoptspyderbattmpbat1648488918mmuid62vqs2LGqooGA1ECM66dqV2PKTl141s",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinmkdirprun",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinlnshostrunk3ssockrunsock",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155bootstrapshbootstrapsh",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinmkdirphostoptspyderbat",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinlnshostoptopt",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinmkdirpvarrun",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinmvoptexistingopt",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinfindhostruninamedockersock",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinhead1",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinfindhostruninamesock",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinlnshostoptopt",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinmkdirpvarrun",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinfindhostvarinamesock",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinhead1",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinfindhostruninamedockersock",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinfindhostruninamecriosock",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinfindhostruninamedockershimsock",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbincurlretry5httpsorcspyderbatcomv1regbL6Ns0xJY0MDliFn0xqXscript",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinwhoami",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinuname",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbingrepCentOSrelease6etcsystemrelease",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinunamem",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinunamea",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinloggerNanoAgentinstallplatformdetectedUbuntudebian220460x8664",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinmkdirpoptspyderbatetc",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinmkdirpoptspyderbattmp",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinmkdirpoptspyderbatbin",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod700optspyderbatetc",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod600optspyderbattmp",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinchmod600optspyderbatbin",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrbintputTscreen256colorslhs",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrbintputTscreenslhs",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrbintty",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrlocalbinkubectlcompletionbash",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrbinreadlinkproc1exe",
	"188bf711eb6446b27b0ea3090fb7cb5ed3b08fcb640be5203c31f4cb359cf283usrbinshckubectlproxydisablefilter||true",
	"188bf711eb6446b27b0ea3090fb7cb5ed3b08fcb640be5203c31f4cb359cf283usrlocalbinkubectlproxydisablefilter",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrlocalbinhelmrepoaddnanoagenthttpsspyderbatgithubionanoagenthelm",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrlocalbinhelmrepoupdate",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrlocalbinhelminstallnanoagentnanoagentnanoagentsetnanoagentagentRegistrationCodebL6Ns0xJY0MDliFn0xqXsetnanoagentorcurlhttpsorcspyderbatcomnamespacespyderbatcreatenamespacesetCLUSTERNAMErke11",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13abinsh",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbingrepLinux",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinawkF{print12}",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinunamer",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161ebootstrapshbootstrapsh",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinmkdirphostoptspyderbat",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinmvoptexistingopt",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinlnshostoptopt",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinmkdirpvarrun",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinfindhostruninamedockersock",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrbinhead1",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13abootstrapshbootstrapsh",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinmkdirphostoptspyderbat",
	"7b6356e9d51649be367ba963a11f5bc57b4ca58ea40dcde2b5c25864072af13ausrbinmvoptexistingopt",
	"0c34f800922e120ea377704d4df099325ed9199d10df58637b21e7a58ce7cddfusrbingitjobgitjobimageranchergitjobv0196",
	"b1fb0383ed0e26cf732daab3421402b164c085389d61a8f60fb6dde35526c686usrbinwebhook",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2ranchercharts4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721resetharddeb8604cf947f508db0ab71a58cc7fe43f6e7592",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2ranchercharts4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721resethardHEAD",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2ranchercharts4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721revparseHEAD",
	"09a1032d43d911649276ef1d6f5f17f67e6600d08cb78ab6704fc5c4b986a52fusrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2ranchercharts4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721resetharddeb8604cf947f508db0ab71a58cc7fe43f6e7592",
	"29aa2f08db8b1663c368f4e2f2eff3b3ed016e1c56ef6ea180969350608d97b3usrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2ranchercharts4b40cac650031b74776e87c1a726b0484d0877c3ec137da0872547ff9b73a721resetharddeb8604cf947f508db0ab71a58cc7fe43f6e7592",
	"29aa2f08db8b1663c368f4e2f2eff3b3ed016e1c56ef6ea180969350608d97b3usrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherrke2charts675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9resethard98e221e79a5c8461c48059d4e819b9325542cad7",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherpartnercharts8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974resethard5b498cba1c410428f26445b061a1a7a5741b9aa4",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherpartnercharts8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974resethardHEAD",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherpartnercharts8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974revparseHEAD",
	"29aa2f08db8b1663c368f4e2f2eff3b3ed016e1c56ef6ea180969350608d97b3usrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherpartnercharts8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974resethard5b498cba1c410428f26445b061a1a7a5741b9aa4",
	"09a1032d43d911649276ef1d6f5f17f67e6600d08cb78ab6704fc5c4b986a52fusrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherpartnercharts8f17acdce9bffd6e05a58a3798840e408c4ea71783381ecd2e9af30baad65974resethard5b498cba1c410428f26445b061a1a7a5741b9aa4",
	"09a1032d43d911649276ef1d6f5f17f67e6600d08cb78ab6704fc5c4b986a52fusrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherrke2charts675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9resethard98e221e79a5c8461c48059d4e819b9325542cad7",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherrke2charts675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9resethard98e221e79a5c8461c48059d4e819b9325542cad7",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherrke2charts675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9resethardHEAD",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitccredentialhelperbinshcechopasswordGITPASSWORDCvarlibrancherdatalocalcatalogsv2rancherrke2charts675f1b63a0a83905972dcab2794479ed599a6f41b86cd6193d69472d0fa889c9revparseHEAD",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitCmanagementstatecatalogcachea67038d6110101d84b823470950db15b8ab6c06fe4e8895bcae7a337c1a8990drevparseHEAD",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitCmanagementstatecatalogcachef341cfdfa521a9aa2b993cb34b26bb91b2d173ef1a7df8d41b8921b0e4f82788revparseHEAD",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbingitCmanagementstatecatalogcache380859f1003fe7603cddc6c15b34b7263f1f0deaa92ddcde465811d032ee7078revparseHEAD",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrlocalbinwelcomeusrlocalbinwelcome",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrbinbash",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrbinreadlinkproc25exe",
	"fbc6f8414b640290e95375f9ed24ede7bad46592ae2eb66c0cbf2326596904fdusrbintpuths",
	"22c10a219637fa34df2d2550fcc2e956faea6933f418e38e9188cd795cbba88busrbinrancherhttplistenport80httpslistenport443auditlogpathvarlogauditlograncherapiauditlogauditlevel1auditlogmaxage1auditlogmaxbackup1auditlogmaxsize100nocacertshttplistenport80httpslistenport443addlocaltrue",
	"29aa2f08db8b1663c368f4e2f2eff3b3ed016e1c56ef6ea180969350608d97b3usrbinrancherhttplistenport80httpslistenport443auditlogpathvarlogauditlograncherapiauditlogauditlevel1auditlogmaxage1auditlogmaxbackup1auditlogmaxsize100nocacertshttplistenport80httpslistenport443addlocaltrue",
	"ade5e497ca26c171ea29730be255cd66461f8007901ff5577dc1efe86ec6e863optredpandalibexecrpkclusterhealth",
	"09a1032d43d911649276ef1d6f5f17f67e6600d08cb78ab6704fc5c4b986a52fusrbinrancherhttplistenport80httpslistenport443auditlogpathvarlogauditlograncherapiauditlogauditlevel1auditlogmaxage1auditlogmaxbackup1auditlogmaxsize100nocacertshttplistenport80httpslistenport443addlocaltrue",
	"e675d30264ce91ca85ef1df2c0bc0ee192b80bdce571d59c13d96bef194555causrbinfleetcontrollerdisablebootstrap",
	"ade5e497ca26c171ea29730be255cd66461f8007901ff5577dc1efe86ec6e863usrbincurlsilentfailkhttpredpandasrc0redpandasrcredpandasvcclusterlocal9644v1statusready",
	"ade5e497ca26c171ea29730be255cd66461f8007901ff5577dc1efe86ec6e863usrbinbashusrbinrpkclusterhealth",
	"ade5e497ca26c171ea29730be255cd66461f8007901ff5577dc1efe86ec6e863binshccurlsilentfailkhttp{SERVICENAME}redpandasrcredpandasvcclusterlocal9644v1statusready",
	"b9b9780cb8517df918d3d5a1cd72b52f6ff49f7b17251767a6ccc5ee7c70161eusrsbiniproute",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrbinunamesrio",
	"2030a2e446c7191353e920cb44da32a7309c0557c233168dd410e33c51ecd155usrsbiniproute",
}

var keysSet map[string]struct{}

func doTransform(e transform.WriteEvent, w transform.RecordWriter) error {
	// Unmarshal the incoming message into a map
	record := e.Record()
	if strings.Contains(string(record.Value), "/var/lib/rancher-data/local-catalogs/v2/rancher") {
		return nil
	}

	var incomingMessage map[string]interface{}
	err := json.Unmarshal(e.Record().Value, &incomingMessage)
	if err != nil {
		return err
	}

	// Extract 3 fields from the JSON and concat them as key
	key := createKey(incomingMessage)

	// Check if the key is in the CSV keys
	if _, ok := keysSet[key]; !ok {
		// If the key is not in the CSV keys, write the message

		// Marshal the result back to JSON
		jsonData, err := json.Marshal(incomingMessage)
		if err != nil {
			return err
		}

		// Create a new record with the JSON data
		record := &transform.Record{
			Key:       []byte(key),
			Value:     jsonData,
			Offset:    e.Record().Offset,
			Timestamp: e.Record().Timestamp,
			Headers:   e.Record().Headers,
		}

		// Write the record to the destination topic
		err = w.Write(*record)
		if err != nil {
			return err
		}
	}

	return nil
}
