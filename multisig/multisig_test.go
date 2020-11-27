package multisig

import (
	"github.com/Zilliqa/gozilliqa-sdk/util"
	"strings"
	"testing"
)

func TestMultiVerify(t *testing.T) {
	msg := "67C6697351FF4AEC29CDBAABF2FBE3467CC254F81BE8E78D765A2E63339FC99A66320DB73158A35A255D051758E95ED4ABB2CDC69BB454110E827441213DDC8770E93EA141E1FC673E017E97EADC6B968F385C2AECB03BFB32AF3C54EC18DB5C021AFE43FBFAAA3AFB29D1E6053C7C9475D8BE6189F95CBBA8990F95B1EBF1B305EFF700E9A13AE5CA0BCBD0484764BD1F231EA81C7B64C514735AC55E4B79633B706424119E09DCAAD4ACF21B10AF3B33CDE3504847155CBB6F2219BA9B7DF50BE11A1C7F23F829F8A41B13B5CA4EE8983238E0794D3D34BC5F4E77FACB6C05AC86212BAA1A55A2BE70B5733B045CD33694B3AFE2F0E49E4F321549FD824EA90870D4B28A2954489A0ABCD50E18A844AC5BF38E4CD72D9B0942E506C433AFCDA3847F2DADD47647DE321CEC4AC430F62023856CFBB20704F4EC0BB920BA86C33E05F1ECD96733B79950A3E314D3D934F75EA0F210A8F6059401BEB4BC4478FA4969E623D01ADA696A7E4C7E5125B34884533A94FB319990325744EE9BBCE9E525CF08F5E9E25E5360AAD2B2D085FA54D835E8D466826498D9A8877565705A8A3F62802944DE7CA5894E5759D351ADAC869580EC17E485F18C0C66F17CC07CBB22FCE466DA610B63AF62BC83B4692F3AFFAF271693AC071FB86D11342D8DEF4F89D4B66335C1C7E4248367D8ED9612EC453902D8E50AF89D7709D1A596C1F41F95AA82CA6C49AE90CD1668BAAC7AA6F2B4A8CA99B2C2372ACB08CF61C9C3805E6E0328DA4CD76A19EDD2D3994C798B0022569AD418D1FEE4D9CD45A391C601FFC92AD91501432FEE150287617C13629E69FC7281CD7165A63EAB49CF714BCE3A75A74F76EA7E64FF81EB61FDFEC39B67BF0DE98C7E4E32BDF97C8C6AC75BA43C02F4B2ED7216ECF3014DF000108B67CF99505B179F8ED4980A6103D1BCA70DBE9BBFAB0ED59801D6E5F2D6F67D3EC5168E212E2DAF02C6B963C98A1F7097DE0C56891A2B211B01070DD8FD8B16C2A1A4E3CFD292D2984B3561D555D16C33DDC2BCF7EDDE13EFE520C7E2ABDDA44D81881C531AEEEB66244C3B791EA8ACFB6A68F3584606472B260E0DD2EBB21F6C3A3BC0542AABBA4EF8F6C7169E731108DB0460220AA74D31B55B03A00D220D475DCD9B877856D5704C9C86EA0F98F2EB9C530DA7FA5AD8B0B5DB50C2FD5D095A2AA5E2A3FBB71347549A316332234ECE765B7571B64D216B28712E25CF3780F9DC629CD719B01E6D4A4FD17C731F4AE97BC05A310D7B9C36EDCA5BBC02DBB5DE3D52B65702D4C44C2495C897B5128030D2DB61E056FD1643C871FFCA4DB5A88A075EE10933A655573B1DEEF02F6E20024981E2A07FF8E34769E311B698B9419F1822A84BC8FDA2041A90F449FE154B48962DE81525CB5C8FAE6D45462786E53FA98D8A718A2C75A4BC6AEEBA7F39021567EA"
	signature := "5D4D7C1927FC21EB5CA071568225D4736F1564418B68806A81574D8A8774BFFBEE5025D422E4006448A3C937B36DEE98CF75F66010CBCF537C30B59B383B3363"
	pubKey := "030C1C24EAA7C656D1DD7AF98FC3DD4D3FBB441BAB9D83D42302E6D0E43F6B1C4D"
	r := util.DecodeHex(signature[0:64])
	s := util.DecodeHex(signature[64:128])
	if !MultiVerify(util.DecodeHex(pubKey), util.DecodeHex(msg), r,s) {
		t.Fail()
	}

	msg = "67C6697351FF4AEC29CDBAABF2FBE3467CC254F81BE8E78D765A2E63339FC99A66320DB73158A35A255D051758E95ED4ABB2CDC69BB454110E827441213DDC8770E93EA141E1FC673E017E97EADC6B968F385C2AECB03BFB32AF3C54EC18DB5C021AFE43FBFAAA3AFB29D1E6053C7C9475D8BE6189F95CBBA8990F95B1EBF1B305EFF700E9A13AE5CA0BCBD0484764BD1F231EA81C7B64C514735AC55E4B79633B706424119E09DCAAD4ACF21B10AF3B33CDE3504847155CBB6F2219BA9B7DF50BE11A1C7F23F829F8A41B13B5CA4EE8983238E0794D3D34BC5F4E77FACB6C05AC86212BAA1A55A2BE70B5733B045CD33694B3AFE2F0E49E4F321549FD824EA90870D4B28A2954489A0ABCD50E18A844AC5BF38E4CD72D9B0942E506C433AFCDA3847F2DADD47647DE321CEC4AC430F62023856CFBB20704F4EC0BB920BA86C33E05F1ECD96733B79950A3E314D3D934F75EA0F210A8F6059401BEB4BC4478FA4969E623D01ADA696A7E4C7E5125B34884533A94FB319990325744EE9BBCE9E525CF08F5E9E25E5360AAD2B2D085FA54D835E8D466826498D9A8877565705A8A3F62802944DE7CA5894E5759D351ADAC869580EC17E485F18C0C66F17CC07CBB22FCE466DA610B63AF62BC83B4692F3AFFAF271693AC071FB86D11342D8DEF4F89D4B66335C1C7E4248367D8ED9612EC453902D8E50AF89D7709D1A596C1F41F95AA82CA6C49AE90CD1668BAAC7AA6F2B4A8CA99B2C2372ACB08CF61C9C3805E6E0328DA4CD76A19EDD2D3994C798B0022569AD418D1FEE4D9CD45A391C601FFC92AD91501432FEE150287617C13629E69FC7281CD7165A63EAB49CF714BCE3A75A74F76EA7E64FF81EB61FDFEC39B67BF0DE98C7E4E32BDF97C8C6AC75BA43C02F4B2ED7216ECF3014DF000108B67CF99505B179F8ED4980A6103D1BCA70DBE9BBFAB0ED59801D6E5F2D6F67D3EC5168E212E2DAF02C6B963C98A1F7097DE0C56891A2B211B01070DD8FD8B16C2A1A4E3CFD292D2984B3561D555D16C33DDC2BCF7EDDE13EFE520C7E2ABDDA44D81881C531AEEEB66244C3B791EA8ACFB6A68F3584606472B260E0DD2EBB21F6C3A3BC0542AABBA4EF8F6C7169E731108DB0460220AA74D31B55B03A00D220D475DCD9B877856D5704C9C86EA0F98F2EB9C530DA7FA5AD8B0B5DB50C2FD5D095A2AA5E2A3FBB71347549A316332234ECE765B7571B64D216B28712E25CF3780F9DC629CD719B01E6D4A4FD17C731F4AE97BC05A310D7B9C36EDCA5BBC02DBB5DE3D52B65702D4C44C2495C897B5128030D2DB61E056FD1643C871FFCA4DB5A88A075EE10933A655573B1DEEF02F6E20024981E2A07FF8E34769E311B698B9419F1822A84BC8FDA2041A90F449FE154B48962DE81525CB5C8FAE6D45462786E53FA98D8A718A2C75A4BC6AEEBA7F39021567EA2B8CB6871B64F561AB1CE7905B901EE502A811774DCDE13B8760748A76DB74A1682A28838F1DE43A39CCCA945CE8795E918AD6DE57B719DF188D698E69DD2FD1085754977539D1AE059B436184BCC0154796F39E4D0C7D6599E6F302C422D3CC7A2863EF61349D66CFE0C7539D8768E41D5B826B6700D001E6C403AAE6D77660FFD94F600DEDC6DDCD8D306A15994E32F4D19D5CD16E5DB73260621837D87936B2C896BFB55C9C83EACDEDFF663C315A0DCFB6DE3D13956F74F787ABD000E282C978417ED5DE01BFABEFBE112BEF6B38BE2216FB35AB6AA9A3F25573F237F5BBAF363A84143B43BF2A01D055F13C8DAF5EA3AB934F153DF2079265FAC95AB57890EFFDA52B4064554235AB337138E2CFDC8D622BA39F1DAA3182A4FADC5A736C49701174B076CAF2AB75251CAD08EB89954DB438EDD1E31E5387192FE18C9C2BFCAD9FAC23699FCEDEC4EA8CCCD5156223CA9A109B7D2EEF05471EE6D3BA11CF68B17C8B1A1B5AF9DF4485AC1A9A0E3D64A84D00267BEF2BC30D1196C8236630D4E2BBEEFD15E7DC5A6C88740796B16B3FFE6B65795A903C68A1D330C43960981B1B8718316EF48BDB7DFFE213B04D52AEB9B7271347647BE937ABAD700B468B27CDA3583B97E31614E2F82892467A40FF32671279CB8E62023910724556FD6C23A0C45E38A7754C896D741BB3EF5BB221C2C59A8E53FD908C0D03D163003D86A101E4D9A8592531C79A4C7A89A72DAA6AF244F8454188D14E8BA3B18CE0372DE21C068A752BBC3CC508B74EB0E4F81AD63D121B7E9AECCD268F7EB270B6DF52D2E5DC47109884D6A13B24511F1D6BF55A7D10D817FCA53D8C24EFFCDACE4EACB32AF3C4C3779A64B2BEB5D1DB20C6359DD60EB4D3B3F25FD7E15BB1B0A95D63D3512796C8C1FA7B80AF4C5BCF13916CE99F21BC52131B2AF476DBA41F3908F38A2F8952F184CD71331ACC032D5D6F16FC90D34FA3EE799865543C848D44771774016A658537D6B851A2BB7E002B95FCBB684B5F56C4F7BB193340A678B7BEECB828513D5F27F6B1C9B12FC9DCC4C6982C11F783D6EE3EEF217E9599365385EE7BD62CDBFD228CC7D3BB90B0805648AC683F2F3E2D6E2D4EECC2E822166D1191443D6C415FF80832B499E234EF2AE057691095967EC2E56A85CD8D9B3A9E2C7EDB99C03A91C86C45614F7951795AA8E36A3E79E8005E52852BDF20667D4DE458E6A492776DFFBDCE4E361FC790C8AAFA0624E20682358CAE14AC1492F9F8EADF9D7D570A7C14D8CA4AF891DBC03CD5C660B8CCE2ED58900105A493FE9D7EDE3AFB354477491C419314D26ED40E449A6EFC6751E9BFE1EAC4867EC323FCA15DF7D6A16E1FBDAFB2D28121A6906541FE61A84F4A6731342CB7B2EFDAAE9037A566D8138595C2376744580ED4BD4FD21EF722685E539D8A0A4F79E4FE091BA36FF3"
	signature = "B0D8AC5F97C9B5FE7200849709A9E8E5B08B5CD23E3CA4778EA5C9978264421257D7A1CF0A32A78DBC4404C5710B781DCFCABC9167643152F523C2C0AAE46C2B"
	pubKey = "030618BBB66FCF69E91A21639FE0962F1CE390093EE67D1D10F54C8C37947051BC"
	r = util.DecodeHex(signature[0:64])
	s = util.DecodeHex(signature[64:128])
	if !MultiVerify(util.DecodeHex(pubKey), util.DecodeHex(msg), r,s) {
		t.Fail()
	}
}

func TestAggregatedPubKey(t *testing.T) {
	// private key 1 44F865C7E5DC389CD8797FB219D1288C53AFFCED13B4FED7F04F1C7AD940F46A
	// private key 2 F981B36B10072FA2B873FB279B872D06B036DBCDA5BF7058AD2CBD927AFF2DBA
	// private key 3 AF8892AAA47CF0E12389E51660146481BF52729025DFABF9E269E53D9247F785
	pubKey1 := util.DecodeHex("0218A720B7EBCC45E7419E353A0782B0A7C33726D72361B978E92B96A669BF3867")
	pubKey2 := util.DecodeHex("029637AB898FD03535B9BB51BE7FB7F48F6E8481E213A000EF5D2BC04A378E7D7B")
	pubKey3 := util.DecodeHex("0357496874D1C65EE90F81BA2215E8881295D9AF5D972C9CCC0E79610632811F3E")
	pubKeys := [][]byte{pubKey1, pubKey2, pubKey3}
	aggregatedPubKey, err := AggregatedPubKey(pubKeys)
	pubKey := util.EncodeHex(aggregatedPubKey)
	if err != nil {
		t.Error(err)
	}
	// the target string is generated by cpp side
	if pubKey != strings.ToLower("02A6EFF003EB4E92EC46930BD2A994BA4B085B43AE70CBBA65AE77168B1807049F") {
		t.Error("generate aggregated public key failed")
	}
}
