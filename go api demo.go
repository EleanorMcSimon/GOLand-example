package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var sendTo = "blank"
var uuid = "blank"
var sellRegion = "blank"
var selleruid = "blank"

type sort struct {
	Direction string `json:"direction"`
	Property  string `json:"property"`
}
type and struct {
	Property string   `json:"property"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}
type page struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
type device struct {
	Filter struct {
		And []and `json:"and"`
	} `json:"filter"`
	Sort    []sort `json:"sort"`
	pagegen page   `json:"pagination"`
}
type connection struct {
	Type       string `json:"type"`
	Name       string `json:"name"`
	Bandwidth  int    `json:"bandwidth"`
	Redundancy struct {
		Priority string `json:"priority"`
		Group    string `json:"group"`
	} `json:"redundancy"`
	ASide struct {
		AccessPoint struct {
			Type         string `json:"type"`
			LinkProtocol struct {
				Type    string `json:"type"`
				VlanTag int    `json:"vlanTag"`
			} `json:"linkProtocol"`
			VirtualDevice struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			} `json:"virtualDevice"`
		} `json:"accessPoint"`
	} `json:"aSide"`
	ZSide struct {
		AccessPoint struct {
			Type    string `json:"type"`
			Profile struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			} `json:"profile"`
			Location struct {
				MetroCode string `json:"metroCode"`
			} `json:"location"`
			AuthenticationKey string `json:"authenticationKey"`
			SellerRegion      string `json:"sellerRegion"`
		} `json:"accessPoint"`
	} `json:"zSide"`
	Notifications []notifications `json:"notifications"`
}
type notifications struct {
	Type   string   `json:"type"`
	Emails []string `json:"emails"`
}
type readback struct {
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Sort []struct {
		Direction string `json:"direction"`
		Property  string `json:"property"`
	} `json:"sort"`
	Data []struct {
		Type   string `json:"type"`
		Href   string `json:"href"`
		Uuid   string `json:"uuid"`
		Name   string `json:"name"`
		State  string `json:"state"`
		Change struct {
			Uuid            string    `json:"uuid"`
			Type            string    `json:"type"`
			Status          string    `json:"status"`
			CreatedDateTime time.Time `json:"createdDateTime"`
			UpdatedDateTime time.Time `json:"updatedDateTime"`
			Data            struct {
				Op    string `json:"op"`
				Path  string `json:"path"`
				Value string `json:"value"`
			} `json:"data"`
		} `json:"change,omitempty"`
		Operation struct {
			ProviderStatus string `json:"providerStatus"`
			EquinixStatus  string `json:"equinixStatus"`
			Errors         []struct {
				ErrorCode    string `json:"errorCode"`
				ErrorMessage string `json:"errorMessage"`
			} `json:"errors,omitempty"`
		} `json:"operation"`
		Order struct {
			BillingTier string `json:"billingTier"`
		} `json:"order"`
		Notifications []struct {
			Type   string   `json:"type"`
			Emails []string `json:"emails"`
		} `json:"notifications"`
		Account struct {
			AccountNumber int    `json:"accountNumber"`
			OrgId         int    `json:"orgId"`
			GlobalOrgId   string `json:"globalOrgId"`
		} `json:"account"`
		ChangeLog struct {
			CreatedBy       string    `json:"createdBy"`
			CreatedDateTime time.Time `json:"createdDateTime"`
			UpdatedBy       string    `json:"updatedBy"`
			UpdatedDateTime time.Time `json:"updatedDateTime"`
			DeletedBy       string    `json:"deletedBy,omitempty"`
			DeletedDateTime time.Time `json:"deletedDateTime,omitempty"`
		} `json:"changeLog"`
		Bandwidth  int `json:"bandwidth"`
		Redundancy struct {
			Group    string `json:"group"`
			Priority string `json:"priority"`
		} `json:"redundancy"`
		IsRemote  bool   `json:"isRemote"`
		Direction string `json:"direction"`
		ASide     struct {
			AccessPoint struct {
				Interface struct {
					Uuid string `json:"uuid"`
					Id   int    `json:"id"`
					Type string `json:"type"`
				} `json:"interface,omitempty"`
				Type    string `json:"type"`
				Account struct {
					AccountNumber    int    `json:"accountNumber"`
					OrganizationName string `json:"organizationName"`
				} `json:"account"`
				Location struct {
					Region    string `json:"region"`
					MetroName string `json:"metroName"`
					MetroCode string `json:"metroCode"`
				} `json:"location"`
				LinkProtocol struct {
					Type    string `json:"type"`
					VlanTag int    `json:"vlanTag"`
				} `json:"linkProtocol"`
				VirtualDevice struct {
					Href string `json:"href"`
					Uuid string `json:"uuid"`
					Name string `json:"name"`
					Type string `json:"type"`
				} `json:"virtualDevice,omitempty"`
			} `json:"accessPoint"`
		} `json:"aSide"`
		ZSide struct {
			AccessPoint struct {
				Type    string `json:"type"`
				Account struct {
					OrganizationName string `json:"organizationName"`
					AccountNumber    int    `json:"accountNumber,omitempty"`
				} `json:"account"`
				Location struct {
					Region    string `json:"region"`
					MetroName string `json:"metroName,omitempty"`
					MetroCode string `json:"metroCode,omitempty"`
				} `json:"location"`
				Port struct {
					Href       string `json:"href"`
					Uuid       string `json:"uuid"`
					Name       string `json:"name"`
					Redundancy struct {
						Priority string `json:"priority"`
					} `json:"redundancy"`
				} `json:"port,omitempty"`
				Profile struct {
					Href string `json:"href"`
					Type string `json:"type"`
					Name string `json:"name"`
					Uuid string `json:"uuid"`
				} `json:"profile,omitempty"`
				LinkProtocol struct {
					Type    string `json:"type"`
					VlanTag int    `json:"vlanTag,omitempty"`
				} `json:"linkProtocol,omitempty"`
				SellerRegion         string `json:"sellerRegion,omitempty"`
				AuthenticationKey    string `json:"authenticationKey,omitempty"`
				ProviderConnectionId string `json:"providerConnectionId,omitempty"`
				Network              struct {
					Uuid  string `json:"uuid"`
					Name  string `json:"name"`
					Type  string `json:"type"`
					Scope string `json:"scope"`
				} `json:"network,omitempty"`
			} `json:"accessPoint"`
			ServiceToken struct {
				Href    string `json:"href"`
				Uuid    string `json:"uuid"`
				Account struct {
					OrgId int `json:"orgId"`
				} `json:"account"`
			} `json:"serviceToken,omitempty"`
		} `json:"zSide"`
		AdditionalInfo []interface{} `json:"additionalInfo,omitempty"`
		Project        struct {
			ProjectId string `json:"projectId,omitempty"`
		} `json:"project"`
	} `json:"data"`
}

type resposeAftermaking struct {
	Type      string `json:"type"`
	Href      string `json:"href"`
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Operation struct {
		ProviderStatus string `json:"providerStatus"`
		EquinixStatus  string `json:"equinixStatus"`
	} `json:"operation"`
	Order struct {
		BillingTier string `json:"billingTier"`
	} `json:"order"`
	Notifications []struct {
		Type   string   `json:"type"`
		Emails []string `json:"emails"`
	} `json:"notifications"`
	ChangeLog struct {
		CreatedBy         string    `json:"createdBy"`
		CreatedByFullName string    `json:"createdByFullName"`
		CreatedByEmail    string    `json:"createdByEmail"`
		CreatedDateTime   time.Time `json:"createdDateTime"`
		UpdatedBy         string    `json:"updatedBy"`
		UpdatedByFullName string    `json:"updatedByFullName"`
		UpdatedByEmail    string    `json:"updatedByEmail"`
		UpdatedDateTime   time.Time `json:"updatedDateTime"`
	} `json:"changeLog"`
	Bandwidth  int `json:"bandwidth"`
	Redundancy struct {
		Group    string `json:"group"`
		Priority string `json:"priority"`
	} `json:"redundancy"`
	ASide struct {
		AccessPoint struct {
			Interface struct {
				Uuid string `json:"uuid"`
				Id   int    `json:"id"`
				Type string `json:"type"`
			} `json:"interface"`
			Location struct {
				MetroCode string `json:"metroCode"`
			} `json:"location"`
			LinkProtocol struct {
				Type    string `json:"type"`
				VlanTag int    `json:"vlanTag"`
			} `json:"linkProtocol"`
			VirtualDevice struct {
				Uuid string `json:"uuid"`
				Type string `json:"type"`
			} `json:"virtualDevice"`
		} `json:"accessPoint"`
	} `json:"aSide"`
	ZSide struct {
		AccessPoint struct {
			Location struct {
				MetroCode string `json:"metroCode"`
			} `json:"location"`
			Profile struct {
				Href string `json:"href"`
				Type string `json:"type"`
				Name string `json:"name"`
				Uuid string `json:"uuid"`
			} `json:"profile"`
			SellerRegion      string `json:"sellerRegion"`
			AuthenticationKey string `json:"authenticationKey"`
		} `json:"accessPoint"`
	} `json:"zSide"`
}
type filter struct {
	Filter struct {
		Property string   `json:"property"`
		Operator string   `json:"operator"`
		Values   []string `json:"values"`
	} `json:"filter"`
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
	} `json:"pagination"`
	Sort []sort `json:"sort"`
}
type rep struct {
	Data []struct {
		State   string `json:"state"`
		Account struct {
			OrganizationName       string `json:"organizationName"`
			GlobalOrganizationName string `json:"globalOrganizationName"`
		} `json:"account"`
		Project struct {
		} `json:"project"`
		ChangeLog struct {
			CreatedBy         string    `json:"createdBy"`
			CreatedByFullName string    `json:"createdByFullName"`
			CreatedByEmail    string    `json:"createdByEmail"`
			CreatedDateTime   time.Time `json:"createdDateTime"`
			UpdatedBy         string    `json:"updatedBy"`
			UpdatedByFullName string    `json:"updatedByFullName"`
			UpdatedByEmail    string    `json:"updatedByEmail"`
			UpdatedDateTime   time.Time `json:"updatedDateTime"`
		} `json:"changeLog"`
		Href                   string `json:"href"`
		Type                   string `json:"type"`
		Name                   string `json:"name"`
		Uuid                   string `json:"uuid"`
		Description            string `json:"description"`
		Visibility             string `json:"visibility"`
		AccessPointTypeConfigs []struct {
			Uuid                       string `json:"uuid"`
			Type                       string `json:"type"`
			SupportedBandwidths        []int  `json:"supportedBandwidths"`
			AllowRemoteConnections     bool   `json:"allowRemoteConnections"`
			AllowCustomBandwidth       bool   `json:"allowCustomBandwidth"`
			AllowBandwidthAutoApproval bool   `json:"allowBandwidthAutoApproval"`
			LinkProtocolConfig         struct {
				EncapsulationStrategy string `json:"encapsulationStrategy"`
				ReuseVlanSTag         bool   `json:"reuseVlanSTag"`
				Encapsulation         string `json:"encapsulation"`
			} `json:"linkProtocolConfig"`
			EnableAutoGenerateServiceKey bool `json:"enableAutoGenerateServiceKey"`
			ConnectionRedundancyRequired bool `json:"connectionRedundancyRequired"`
			ApiConfig                    struct {
				ApiAvailable     bool   `json:"apiAvailable"`
				IntegrationId    string `json:"integrationId"`
				BandwidthFromApi bool   `json:"bandwidthFromApi"`
			} `json:"apiConfig"`
			ConnectionLabel   string `json:"connectionLabel"`
			AuthenticationKey struct {
				Required bool   `json:"required"`
				Label    string `json:"label"`
			} `json:"authenticationKey"`
			Metadata struct {
				RegEx              string `json:"regEx"`
				RegExMsg           string `json:"regExMsg"`
				GlobalOrganization string `json:"globalOrganization"`
				LimitAuthKeyConn   bool   `json:"limitAuthKeyConn"`
				AllowVcMigration   bool   `json:"allowVcMigration"`
				ConnectionEditable bool   `json:"connectionEditable"`
			} `json:"metadata"`
		} `json:"accessPointTypeConfigs"`
		MarketingInfo struct {
			Promotion bool `json:"promotion"`
		} `json:"marketingInfo"`
		Metros []struct {
			Code          string   `json:"code"`
			Name          string   `json:"name"`
			Ibxs          []string `json:"ibxs"`
			DisplayName   string   `json:"displayName"`
			SellerRegions struct {
				ApSoutheast1 string `json:"ap-southeast-1,omitempty"`
				ApEast1      string `json:"ap-east-1,omitempty"`
				SaEast1      string `json:"sa-east-1,omitempty"`
				ApNortheast1 string `json:"ap-northeast-1,omitempty"`
				EuCentral1   string `json:"eu-central-1,omitempty"`
				UsEast1      string `json:"us-east-1,omitempty"`
				UsWest2      string `json:"us-west-2,omitempty"`
				EuWest2      string `json:"eu-west-2,omitempty"`
				EuWest1      string `json:"eu-west-1,omitempty"`
				ApSouth1     string `json:"ap-south-1,omitempty"`
				EuNorth1     string `json:"eu-north-1,omitempty"`
				ApSoutheast2 string `json:"ap-southeast-2,omitempty"`
				ApSoutheast4 string `json:"ap-southeast-4,omitempty"`
				ApNortheast2 string `json:"ap-northeast-2,omitempty"`
				EuSouth1     string `json:"eu-south-1,omitempty"`
				UsWest1      string `json:"us-west-1,omitempty"`
				UsEast2      string `json:"us-east-2,omitempty"`
				CaCentral1   string `json:"ca-central-1,omitempty"`
				EuCentral2   string `json:"eu-central-2,omitempty"`
				EuWest3      string `json:"eu-west-3,omitempty"`
			} `json:"sellerRegions"`
		} `json:"metros"`
		SelfProfile bool `json:"selfProfile"`
	} `json:"data"`
	Pagination struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Total  int `json:"total"`
	} `json:"pagination"`
}

func main() {
	refreschtoken("2IInA8bWggIYPsurWkoXdQIvVaVnnEH7VGaBVcl5lpokvBkb", "KlJUxEYeKNlgvsUsdiAwlFhwDINnNS9GwRkpUMJwRRb2OsUQJqlgcspp8KCaJ5AA")
	getdevis("NY", "VD")
	lookbymetro("NY")
	makeconnection(2, "NY")
}
func makeconnection(times int, metro string) {
	email := []string{"esimon@equinix.com"}
	emails := notifications{
		Type:   "ALL",
		Emails: email,
	}
	note := []notifications{emails}

	for x := 0; x < times; x++ {
		name := "c" + strconv.Itoa(x)
		connection1 := connection{Type: "EVPL_VC", Name: name, Bandwidth: 1000, Redundancy: struct {
			Priority string `json:"priority"`
			Group    string `json:"group"`
		}{"PRIMARY", ""}, ASide: struct {
			AccessPoint struct {
				Type         string `json:"type"`
				LinkProtocol struct {
					Type    string `json:"type"`
					VlanTag int    `json:"vlanTag"`
				} `json:"linkProtocol"`
				VirtualDevice struct {
					Type string `json:"type"`
					Uuid string `json:"uuid"`
				} `json:"virtualDevice"`
			} `json:"accessPoint"`
		}{AccessPoint: struct {
			Type         string `json:"type"`
			LinkProtocol struct {
				Type    string `json:"type"`
				VlanTag int    `json:"vlanTag"`
			} `json:"linkProtocol"`
			VirtualDevice struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			} `json:"virtualDevice"`
		}(struct {
			Type         string
			LinkProtocol struct {
				Type    string `json:"type"`
				VlanTag int    `json:"vlanTag"`
			}
			VirtualDevice struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			}
		}{Type: "VD", LinkProtocol: struct {
			Type    string `json:"type"`
			VlanTag int    `json:"vlanTag"`
		}(struct {
			Type    string
			VlanTag int
		}{Type: "DOT1Q", VlanTag: 1001}), VirtualDevice: struct {
			Type string `json:"type"`
			Uuid string `json:"uuid"`
		}(struct {
			Type string
			Uuid string
		}{Type: "EDGE", Uuid: uuid})})}, ZSide: struct {
			AccessPoint struct {
				Type    string `json:"type"`
				Profile struct {
					Type string `json:"type"`
					Uuid string `json:"uuid"`
				} `json:"profile"`
				Location struct {
					MetroCode string `json:"metroCode"`
				} `json:"location"`
				AuthenticationKey string `json:"authenticationKey"`
				SellerRegion      string `json:"sellerRegion"`
			} `json:"accessPoint"`
		}{AccessPoint: struct {
			Type    string `json:"type"`
			Profile struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			} `json:"profile"`
			Location struct {
				MetroCode string `json:"metroCode"`
			} `json:"location"`
			AuthenticationKey string `json:"authenticationKey"`
			SellerRegion      string `json:"sellerRegion"`
		}(struct {
			Type    string
			Profile struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			}
			Location struct {
				MetroCode string `json:"metroCode"`
			}
			AuthenticationKey string
			SellerRegion      string
		}{Type: "SP", Profile: struct {
			Type string `json:"type"`
			Uuid string `json:"uuid"`
		}(struct {
			Type string
			Uuid string
		}{Type: "L2_PROFILE", Uuid: selleruid}), Location: struct {
			MetroCode string `json:"metroCode"`
		}(struct{ MetroCode string }{MetroCode: metro}), AuthenticationKey: "067257798381", SellerRegion: sellRegion})}, Notifications: note}
		encoded, erro := json.Marshal(connection1)
		os.Stdout.Write(encoded)

		if erro != nil {
			print(erro.Error())
		}
		var bearer = "Bearer " + sendTo //authicatio
		req, erro := http.NewRequest("POST", "https://api.equinix.com/fabric/v4/connections", bytes.NewBuffer(encoded))
		req.Header.Add("Authorization", bearer)
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, erro := client.Do(req)
		print(resp.Status)
		making := resposeAftermaking{}
		bad := json.NewDecoder(resp.Body).Decode(&making)
		if bad != nil {
			print(erro.Error())
		}
		name1 := "cr" + strconv.Itoa(x)
		connection2 := connection{Type: "EVPL_VC", Name: name1, Bandwidth: 1000, Redundancy: struct {
			Priority string `json:"priority"`
			Group    string `json:"group"`
		}{"SECONDARY", making.Redundancy.Group}, ASide: struct {
			AccessPoint struct {
				Type         string `json:"type"`
				LinkProtocol struct {
					Type    string `json:"type"`
					VlanTag int    `json:"vlanTag"`
				} `json:"linkProtocol"`
				VirtualDevice struct {
					Type string `json:"type"`
					Uuid string `json:"uuid"`
				} `json:"virtualDevice"`
			} `json:"accessPoint"`
		}{AccessPoint: struct {
			Type         string `json:"type"`
			LinkProtocol struct {
				Type    string `json:"type"`
				VlanTag int    `json:"vlanTag"`
			} `json:"linkProtocol"`
			VirtualDevice struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			} `json:"virtualDevice"`
		}(struct {
			Type         string
			LinkProtocol struct {
				Type    string `json:"type"`
				VlanTag int    `json:"vlanTag"`
			}
			VirtualDevice struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			}
		}{Type: "VD", LinkProtocol: struct {
			Type    string `json:"type"`
			VlanTag int    `json:"vlanTag"`
		}(struct {
			Type    string
			VlanTag int
		}{Type: "DOT1Q", VlanTag: 1001}), VirtualDevice: struct {
			Type string `json:"type"`
			Uuid string `json:"uuid"`
		}(struct {
			Type string
			Uuid string
		}{Type: "EDGE", Uuid: uuid})})}, ZSide: struct {
			AccessPoint struct {
				Type    string `json:"type"`
				Profile struct {
					Type string `json:"type"`
					Uuid string `json:"uuid"`
				} `json:"profile"`
				Location struct {
					MetroCode string `json:"metroCode"`
				} `json:"location"`
				AuthenticationKey string `json:"authenticationKey"`
				SellerRegion      string `json:"sellerRegion"`
			} `json:"accessPoint"`
		}{AccessPoint: struct {
			Type    string `json:"type"`
			Profile struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			} `json:"profile"`
			Location struct {
				MetroCode string `json:"metroCode"`
			} `json:"location"`
			AuthenticationKey string `json:"authenticationKey"`
			SellerRegion      string `json:"sellerRegion"`
		}(struct {
			Type    string
			Profile struct {
				Type string `json:"type"`
				Uuid string `json:"uuid"`
			}
			Location struct {
				MetroCode string `json:"metroCode"`
			}
			AuthenticationKey string
			SellerRegion      string
		}{Type: "SP", Profile: struct {
			Type string `json:"type"`
			Uuid string `json:"uuid"`
		}(struct {
			Type string
			Uuid string
		}{Type: "L2_PROFILE", Uuid: selleruid}), Location: struct {
			MetroCode string `json:"metroCode"`
		}(struct{ MetroCode string }{MetroCode: metro}), AuthenticationKey: "067257798381", SellerRegion: sellRegion})}, Notifications: note}
		encoded2, erro := json.Marshal(connection2)
		os.Stdout.Write(encoded2)

		req2, erro := http.NewRequest("POST", "https://api.equinix.com/fabric/v4/connections", bytes.NewBuffer(encoded2))
		req2.Header.Add("Authorization", bearer)
		req2.Header.Set("Content-Type", "application/json")

		resp2, erro := client.Do(req2)
		print(resp2.Status)
	}
}
func refreschtoken(id string, clientsecret string) {
	type respose struct { //respose from the server
		AccessToken         string `json:"access_token"`
		TokenTimeout        string `json:"token_timeout"`
		UserName            string `json:"user_name"`
		TokenType           string `json:"token_type"`
		RefreshToken        string `json:"refresh_token"`
		RefreshTokenTimeout string `json:"refresh_token_timeout"`
	}
	type requesttoken struct { // sending to the fabric api
		GrantType    string `json:"grant_type"`
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}
	r := requesttoken{GrantType: "client_credentials", ClientId: id,
		ClientSecret: clientsecret} //sent fields for request
	//b := new(bytes.Buffer)
	encoded, _ := json.Marshal(r) // encoding
	//json.NewEncoder(b).Encode(encoded)

	os.Stdout.Write(encoded)
	var request, erron = http.Post(
		"https://api.equinix.com/oauth2/v1/token",
		"text/json", bytes.NewReader(encoded),
	) // make post request to the api sever to refresh the token

	if erron != nil {
		print(erron.Error())
	}

	var f respose // respose object

	print(request.Status)
	baderr := json.NewDecoder(request.Body).Decode(&f) // decode respose

	// decode to object

	if baderr != nil {
		print(baderr.Error())
	}
	defer request.Body.Close()
	print("done")

	sendTo = f.AccessToken // set access token for feture  all
	print(sendTo)
}
func lookbymetro(metro string) {
	metros := []string{metro}
	s := sort{
		Direction: "DESC",
		Property:  "/changeLog/updatedDateTime",
	}
	sn := []sort{s}
	seachingAws := filter{
		Filter: struct {
			Property string   `json:"property"`
			Operator string   `json:"operator"`
			Values   []string `json:"values"`
		}{"/metros/code", "=", metros},
		Pagination: struct {
			Offset int `json:"offset"`
			Limit  int `json:"limit"`
		}{0, 100000},
		Sort: sn,
	}
	encoded, erro := json.Marshal(seachingAws)
	os.Stdout.Write(encoded)
	if erro != nil {
		print(erro.Error())
	}
	var bearer = "Bearer " + sendTo //authicatio
	req, erro := http.NewRequest("POST", "https://api.equinix.com/fabric/v4/serviceProfiles/search", bytes.NewBuffer(encoded))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, erro := client.Do(req)
	print(resp.Status)
	more := rep{}
	erro = json.NewDecoder(resp.Body).Decode(&more)
	for x := 0; x < len(more.Data); x++ {
		for y := 0; y < len(more.Data[x].Metros); y++ {
			if strings.Compare(more.Data[x].Name, "AWS Direct Connect - High Capacity") == 0 {

				if strings.Compare(more.Data[x].Metros[y].Code, metro) == 0 {

					v := reflect.ValueOf(more.Data[x].Metros[y].SellerRegions)

					values := make([]interface{}, v.NumField())

					for i := 0; i < v.NumField(); i++ {
						values[i] = v.Field(i).String()
						print((v.Field(i).String()))
						if len(v.Field(i).String()) != 0 {
							sellRegion = v.Field(i).String()
						}
						selleruid = more.Data[x].Uuid

					}
				}

			}
		}
	}

}

func getdevis(metro string, div string) {
	key := []string{metro}
	what := []string{div}
	metros := and{
		Property: "/aSide/accessPoint/type",
		Operator: "=",
		Values:   what,
	}
	ty := and{
		Property: "/aSide/accessPoint/location/metroCode",
		Operator: "=",
		Values:   key,
	}

	array := []and{metros, ty}
	sortby := sort{
		Direction: "DESC",
		Property:  "/changeLog/updatedDateTime",
	}
	sorting := []sort{sortby}
	newsearch := device{
		Filter: struct {
			And []and `json:"and"`
		}{And: array},
		Sort: sorting,
		pagegen: page{

			Limit:  40,
			Offset: 0,
		},
	}
	encoded, erro := json.Marshal(newsearch)
	os.Stdout.Write(encoded)
	if erro != nil {
		print(erro.Error())
	}

	var bearer = "Bearer " + sendTo //authicatio
	req, erro := http.NewRequest("POST", "https://api.equinix.com/fabric/v4/connections/search", bytes.NewBuffer(encoded))
	req.Header.Add("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, erro := client.Do(req)
	print(resp.Status)
	hero := readback{}
	erro = json.NewDecoder(resp.Body).Decode(&hero)
	for x := 0; x < len(hero.Data); x++ {
		if strings.Compare(hero.Data[x].Type, "EVPL_VC") == 0 {
			if strings.Compare(hero.Data[x].ASide.AccessPoint.Type, div) == 0 {
				if strings.Compare(hero.Data[x].State, "DEPROVISIONED") != 0 {
					uuid = hero.Data[x].ASide.AccessPoint.VirtualDevice.Uuid
					print(" valid uuid for vds ")
					print(uuid)

				}
			}

		}
	}

}
