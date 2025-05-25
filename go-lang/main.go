package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type WebCategory struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Level int    `json:"level"`
}

type ProductPrice struct {
	SellingPrice        float64 `json:"sellingPrice"`
	BuyingPrice         float64 `json:"buyingPrice"`
	DiscountedPrice     float64 `json:"discountedPrice"`
	DiscountedPriceText string  `json:"discountedPriceText"`
	OriginalPrice       float64 `json:"originalPrice"`
	OriginalPriceText   string  `json:"originalPriceText"`
	SellingPriceText    string  `json:"sellingPriceText"`
	BuyingPriceText     string  `json:"buyingPriceText"`
	Currency            string  `json:"currency"`
	DiscountPercentage  int     `json:"discountPercentage"`
	RRP                 float64 `json:"rrp"`
	ProfitMargin        float64 `json:"profitMargin"`
}

type Product struct {
	ID       int    `json:"id"`
	Url      string `json:"productDetailUrl"`
	Name     string `json:"name"`
	Code     string `json:"productCode"`
	Category struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Hierarchy string `json:"hierarchy"`
	} `json:"category"`
	Seller struct {
		TaxOffice              string `json:"taxOffice"`
		OfficialName           string `json:"officialName"`
		RegisteredEmailAddress string `json:"registeredEmailAddress"`
		RegistrationNumber     string `json:"registrationNumber"`
		TaxNumber              string `json:"taxNumber"`
		Address                string `json:"address"`
		BusinessType           string `json:"businessType"`
		CodEligible            bool   `json:"codEligible"`
	} `json:"sellerInfo"`
	WebCategoryTree []WebCategory `json:"webCategoryTree"`
	RatingScore     struct {
		AverageRating float64 `json:"averageRating"`
		CommentCount  int     `json:"commentCount"`
		TotalCount    int     `json:"totalCount"`
	} `json:"ratingScore"`
	Price    ProductPrice `json:"price"`
	Quantity int          `json:"quantity"`
}

type ProductItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Brand       string `json:"brand"`
	Image       string `json:"image"`
	RatingScore struct {
		AverageRating float64 `json:"averageRating"`
	} `json:"ratingScore"`
	Price struct {
		SellingPrice float64 `json:"sellingPrice"`
	} `json:"price"`
}

type Root struct {
	Data struct {
		Contents []ProductItem `json:"contents"`
	} `json:"data"`
}

func fetchProductDetails(productID int) map[string]interface{} {
	url := fmt.Sprintf("https://apigw.trendyol.com/discovery-sfint-product-service/api/product-detail/?contentId=%d&campaignId=null&storefrontId=36&culture=en-AE", productID) // Replace with real API

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("if-none-match", `W/"2b9d-hQQAkXZ6TfXuSxKb4d/3BO1fF/4"`)
	req.Header.Set("origin", "https://www.trendyol.com")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="135", "Not-A.Brand";v="8", "Chromium";v="135"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
	req.Header.Set("cookie", "platform=web; anonUserId=c88ca0a0-202a-11f0-b282-e524755a5054; OptanonAlertBoxClosed=2025-04-23T10:07:40.439Z; pid=c88ca0a0-202a-11f0-b282-e524755a5054; storefrontId=36; countryCode=AE; language=en; _gcl_au=1.1.140209073.1745402880; _scid=kApt02EUMDwDAF4F9vfx7LsfKVHQYORB; _fbp=fb.1.1745402880754.810322442114590411; _ScCbts=%5B%5D; _pin_unauth=dWlkPU5UZ3hNREptTWpJdE1qZ3hPUzAwT0RFMExXSXlOR010TlRreFpESXdOMlZoWTJVNA; _tt_enable_cookie=1; _ttp=01JSH1WRWS90XBGHC34CSJZ9SE_.tt.1; AbTestingCookies=A_82-B_38-C_43-D_9-E_38-F_99-G_63-H_9-I_13-J_94-K_50-L_55-M_17-N_13-O_80; hvtb=1; VisitCount=1; WebAbTesting=A_16-B_93-C_88-D_50-E_32-F_13-G_69-H_13-I_35-J_51-K_50-L_84-M_21-N_65-O_66-P_10-Q_76-R_88-S_88-T_73-U_72-V_88-W_6-X_53-Y_65-Z_31; msearchAb=ABAdvertSlotPeriod_1-ABDsNlp_2-ABQR_B-ABSearchFETestV1_B-ABBSA_D-ABSuggestionLC_B; AbTesting=pdpAiReviewSummaryUat_B-SFWBAA_V1_B-SFWDBSR_A-SFWDQL_B-SFWDRS_A-SFWDSAOFv2_B-SFWDSFAG_B-SFWDTKV2_A-SSTPRFL_B-STSBuynow_B-STSCouponV2_A-STSImageSocialProof_A-STSRecoRR_B-STSRecoSocialProof_A-WCBsQckFiltTestv2_B-WCOnePageCheckout_B-WEBSFAATest1_A-WebSFAATest2_B-WebSFAATest3_A%7C1745405135%7Cc88ca0a0-202a-11f0-b282-e524755a5054; navbarGenderId=1; guest_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cm46dHJlbmR5b2w6YW5vbmlkIjoiNDMwNmY3M2MyMDQxMTFmMGEwNDNmNjEyMjNiNDc3MGMiLCJyb2xlIjoiYW5vbiIsImF0d3J0bWsiOiI0MzA2ZjczOS0yMDQxLTExZjAtYTA0My1mNjEyMjNiNDc3MGMiLCJhcHBOYW1lIjoidHkiLCJhdWQiOiJzYkF5ell0WCtqaGVMNGlmVld5NXR5TU9MUEpXQnJrYSIsImV4cCI6MTkwMzIwMDUxMSwiaXNzIjoiYXV0aC50cmVuZHlvbC5jb20iLCJuYmYiOjE3NDU0MTI1MTF9.EltSK08NpXAye9_vA86ZAcN_-pIBafYFkFS0uwKe244; csrf-secret=jx65ssufAbKAOyuFhlR442UG; functionalConsent=true; performanceConsent=true; targetingConsent=true; WebRecoTss=collectionRecoVersion%2F1%7CpdpGatewayVersion%2F1%7CsimilarRecoAdsVersion%2F1%7CbasketRecoVersion%2F1%7CsimilarRecoVersion%2F1%7CcompleteTheLookVersion%2F1%7CshopTheLookVersion%2F1%7CcrossRecoAdsVersion%2F1%7CsimilarSameBrandVersion%2F1%7CcrossSameBrandVersion%2F1%7CallInOneRecoVersion%2F1%7CcrossRecoVersion%2F1%7ChomepageVersion%2FfirstComponent%3AinitialNewTest_1.sorter%3AhomepageSorterNewTest_d(M)%7CnavigationSectionVersion%2Fsection%3AazSectionTest_1(M)%7CnavigationSideMenuVersion%2FsideMenu%3AinitialTest_1(M)%7CfirstComponent_V1%2F1%7Csorter_V1%2Fb%7Csection_V1%2F1%7CsideMenu_V1%2F1%7CtopWidgets_V1%2F1; __cf_bm=yAcNpc9.Z023H2FTR6J2dRbKaDczMk24CNVZsuBSnqM-1745517402-1.0.1.1-qjYSPztv9v8GvHS5G0uYTaLZdwoPLOUaz13d12HfxDNO92vJq5WWPX7ZPuSZW4llzjTwiB20ukqB8OOTRtwaJKMxKQA6ZbY2F_GITvRb0yU; _cfuvid=Absmqzycbc8IZ65QvA_J.1CTRkwYDuZ1k_jFWAsDg74-1745517402311-0.0.1.1-604800000; __cflb=04dToYCH9RsdhPpttacYW22gpq3mLXZXuCfT4Kmdad; UserInfo=%7B%22Gender%22%3Anull%2C%22UserTypeStatus%22%3Anull%2C%22ForceSet%22%3Afalse%7D; sid=859166fc-0fcb-478e-a642-7b36856ec13d; _gid=GA1.2.1275501495.1745517404; _dc_gtm_UA-13174585-70=1; ttcsid_CJ5M5PJC77U7DSNBELOG=1745517404716::SmLe8QbNf0ibzXveWLA3.4.1745517406958; ttcsid=1745517404716::bIhYpvqxso6xBLcBmryx.4.1745517406958; tss=firstComponent_V1_1%2Csorter_V1_b%2Csection_V1_1%2CsideMenu_V1_1%2CtopWidgets_V1_1%2CFSA_B%2CProductCardVariantCount_B%2CSuggestionPopular_B%2CRR_2%2CGRRLO_B%2CGRRIn_B%2CVisualCategorySlider_B%2CSuggestionTermActive_B%2CKB_B%2CDGB_B%2CSB_B%2CSuggestion_B%2COFIR_B; _scid_r=rYpt02EUMDwDAF4F9vfx7LsfKVHQYORBse8OXA; _ga=GA1.2.2084360745.1745402880; _uetsid=7c2e4db0213511f0bca4851cb46f5236|1i11fu8|2|fvc|0|1940; _sc_cspv=https%3A%2F%2Ftr.snapchat.com%2Fconfig%2Fcom%2F5df43118-abd2-4cd8-89b9-8cf942d1ee25.js%3Fv%3D3.44.6-2504241707; OptanonConsent=isGpcEnabled=0&datestamp=Thu+Apr+24+2025+22%3A56%3A55+GMT%2B0500+(Pakistan+Standard+Time)&version=202402.1.0&browserGpcFlag=0&isIABGlobal=false&consentId=92cde421-80dd-4b6c-9752-f4bc74338ad3&interactionCount=1&isAnonUser=1&landingPath=NotLandingPage&groups=C0002%3A1%2CC0009%3A1%2CC0007%3A1%2CC0003%3A1%2CC0001%3A1%2CC0004%3A1&hosts=H138%3A1%2CH29%3A1%2CH111%3A1%2CH129%3A1%2CH93%3A1%2CH128%3A1%2CH112%3A1%2CH147%3A1%2CH148%3A1%2CH56%3A1%2CH58%3A1%2CH59%3A1%2CH91%3A1%2CH20%3A1%2CH104%3A1%2CH115%3A1%2CH75%3A1%2CH86%3A1%2CH25%3A1%2CH90%3A1%2CH32%3A1%2CH116%3A1%2CH124%3A1%2CH7%3A1%2CH152%3A1%2CH37%3A1%2CH42%3A1%2CH43%3A1%2CH153%3A1%2CH149%3A1%2CH145%3A1%2CH134%3A1%2CH139%3A1%2CH144%3A1&genVendors=V77%3A1%2CV67%3A1%2CV79%3A1%2CV71%3A1%2CV69%3A1%2CV7%3A1%2CV5%3A1%2CV9%3A1%2CV1%3A1%2CV70%3A1%2CV3%3A1%2CV68%3A1%2CV78%3A1%2CV17%3A1%2CV76%3A1%2CV80%3A1%2CV16%3A1%2CV72%3A1%2CV10%3A1%2CV40%3A1%2C&geolocation=PK%3BPB&AwaitingReconsent=false; cto_bundle=MqRYUl9DNGtMb3BmVDRnNWVnNWJkMTVQcmFHUTIlMkYwdGJSV3NmVTBiWklUM1dsUDQlMkI1M094QjN2QnFtVGNxeXhiWDJzZnc5dm1MckQlMkZYYlNhODAydm1uZkdzMiUyQnNGYnp5TjN1c2lUcndjJTJCdExiRzBjSTRDMXpadXFiNDlBZXhHRnM2eFk5RVdyQ0VwNDA3UTRpbllQaDJrVzElMkJNd1klMkJJRWFaZyUyQnNlZjdRWER6TG5LOVhISG16JTJGUUR0RVFwZDRVSGxza2FoZzFvWXJvbmJWNENJSGtZaGhoY0dBJTNEJTNE; _uetvid=d67945d0202a11f0817e972d9bec397e|19nng9p|1745517415909|2|1|bat.bing.com/p/insights/c/a; _ga_9J2BFGDX1E=GS1.1.1745517404.5.1.1745517462.2.0.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(bodyText, &jsonResponse); err != nil {
		log.Fatal(err)
	}
	return jsonResponse
}

func fetchProductReviews(productID int) map[string]interface{} {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://apigw.trendyol.com/discovery-sfint-social-service/api/review/reviews/358513138?page=1&pageSize=20&storefrontId=36&language=en&countryCode=AE&culture=en-AE&&sellerId=968", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-language", "en-US,en;q=0.9")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://www.trendyol.com")
	req.Header.Set("priority", "u=1, i")
	req.Header.Set("sec-ch-ua", `"Google Chrome";v="135", "Not-A.Brand";v="8", "Chromium";v="135"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
	req.Header.Set("cookie", "platform=web; anonUserId=c88ca0a0-202a-11f0-b282-e524755a5054; OptanonAlertBoxClosed=2025-04-23T10:07:40.439Z; pid=c88ca0a0-202a-11f0-b282-e524755a5054; storefrontId=36; countryCode=AE; language=en; _gcl_au=1.1.140209073.1745402880; _scid=kApt02EUMDwDAF4F9vfx7LsfKVHQYORB; _fbp=fb.1.1745402880754.810322442114590411; _ScCbts=%5B%5D; _pin_unauth=dWlkPU5UZ3hNREptTWpJdE1qZ3hPUzAwT0RFMExXSXlOR010TlRreFpESXdOMlZoWTJVNA; _tt_enable_cookie=1; _ttp=01JSH1WRWS90XBGHC34CSJZ9SE_.tt.1; AbTestingCookies=A_82-B_38-C_43-D_9-E_38-F_99-G_63-H_9-I_13-J_94-K_50-L_55-M_17-N_13-O_80; hvtb=1; VisitCount=1; WebAbTesting=A_16-B_93-C_88-D_50-E_32-F_13-G_69-H_13-I_35-J_51-K_50-L_84-M_21-N_65-O_66-P_10-Q_76-R_88-S_88-T_73-U_72-V_88-W_6-X_53-Y_65-Z_31; msearchAb=ABAdvertSlotPeriod_1-ABDsNlp_2-ABQR_B-ABSearchFETestV1_B-ABBSA_D-ABSuggestionLC_B; AbTesting=pdpAiReviewSummaryUat_B-SFWBAA_V1_B-SFWDBSR_A-SFWDQL_B-SFWDRS_A-SFWDSAOFv2_B-SFWDSFAG_B-SFWDTKV2_A-SSTPRFL_B-STSBuynow_B-STSCouponV2_A-STSImageSocialProof_A-STSRecoRR_B-STSRecoSocialProof_A-WCBsQckFiltTestv2_B-WCOnePageCheckout_B-WEBSFAATest1_A-WebSFAATest2_B-WebSFAATest3_A%7C1745405135%7Cc88ca0a0-202a-11f0-b282-e524755a5054; guest_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1cm46dHJlbmR5b2w6YW5vbmlkIjoiNDMwNmY3M2MyMDQxMTFmMGEwNDNmNjEyMjNiNDc3MGMiLCJyb2xlIjoiYW5vbiIsImF0d3J0bWsiOiI0MzA2ZjczOS0yMDQxLTExZjAtYTA0My1mNjEyMjNiNDc3MGMiLCJhcHBOYW1lIjoidHkiLCJhdWQiOiJzYkF5ell0WCtqaGVMNGlmVld5NXR5TU9MUEpXQnJrYSIsImV4cCI6MTkwMzIwMDUxMSwiaXNzIjoiYXV0aC50cmVuZHlvbC5jb20iLCJuYmYiOjE3NDU0MTI1MTF9.EltSK08NpXAye9_vA86ZAcN_-pIBafYFkFS0uwKe244; _gid=GA1.2.1275501495.1745517404; csrf-secret=FeLW7Ad6PQiFn8-aPCidFu0K; functionalConsent=true; performanceConsent=true; targetingConsent=true; _cfuvid=X1wShUD_X3shAHOFzB5X3FAJOITtCJaCzICtYhN3kzU-1745665465983-0.0.1.1-604800000; UserInfo=%7B%22Gender%22%3Anull%2C%22UserTypeStatus%22%3Anull%2C%22ForceSet%22%3Afalse%7D; sid=935ea0f5-01bd-4876-8e83-d60d1e104772; navbarGenderId=1; cto_bundle=16jy219DNGtMb3BmVDRnNWVnNWJkMTVQcmFLd1d3MFpFeldDd3B0amZOaU9qNTMlMkZ3dGthbUc2NFZiZEhKTmslMkIlMkZmMGh5ZzN4aFIlMkYyVDJ6N043bEw2a0EwdmpQYyUyQkF0eG9wNHVRbTdzMFAxbzVCJTJGeHhBdG12UmF0NmFBd0Vrc2dYJTJCUkpIT1M4cm1LbUxpVEdNMkQxVkptaUFoWjZLOG56OCUyRkZsQlZiMHI4VWJST3pJJTNE; __cf_bm=H7eRlSG1eSDZ3XUtYgl3WZdA1d6JDa0yM5YTOaIJ2ms-1745668172-1.0.1.1-tI6LBBJl0zksqxRouhgM0ZR17OnTz27do_pDFRN1f3RQpZm5_k35g_Hg8nzqCp3iSLWq6QXeQFSdflzh4a.AgGQWXP7qfFRSqgfFG541.yM; __cflb=0pg1Qws9JUp8bBruDeeYR8Qm31HaEK29KHZtpyjb; WebRecoTss=collectionRecoVersion%2F1%7CcrossRecoAdsVersion%2F1%7CsimilarRecoAdsVersion%2F1%7CsimilarSameBrandVersion%2F1%7CcrossSameBrandVersion%2F1%7CallInOneRecoVersion%2F1%7CbasketRecoVersion%2F1%7CcrossRecoVersion%2F1%7CsimilarRecoVersion%2F1%7CcompleteTheLookVersion%2F1%7CshopTheLookVersion%2F1%7CpdpGatewayVersion%2F1%7ChomepageVersion%2FfirstComponent%3AinitialNewTest_1.sorter%3AhomepageSorterNewTest_d(M)%7CnavigationSectionVersion%2Fsection%3AazSectionTest_1(M)%7CnavigationSideMenuVersion%2FsideMenu%3AinitialTest_1(M)%7CfirstComponent_V1%2F1%7Csorter_V1%2Fb%7Csection_V1%2F1%7CsideMenu_V1%2F1%7CtopWidgets_V1%2F1; tss=firstComponent_V1_1%2Csorter_V1_b%2Csection_V1_1%2CsideMenu_V1_1%2CtopWidgets_V1_1%2CFSA_B%2CProductCardVariantCount_B%2CSuggestionPopular_B%2CRR_2%2CGRRLO_B%2CGRRIn_B%2CVisualCategorySlider_B%2CSuggestionTermActive_B%2CKB_B%2CDGB_B%2CSB_B%2CSuggestion_B%2COFIR_B; _dc_gtm_UA-13174585-70=1; OptanonConsent=isGpcEnabled=0&datestamp=Sat+Apr+26+2025+16%3A52%3A42+GMT%2B0500+(Pakistan+Standard+Time)&version=202402.1.0&browserGpcFlag=0&isIABGlobal=false&consentId=92cde421-80dd-4b6c-9752-f4bc74338ad3&interactionCount=1&isAnonUser=1&landingPath=NotLandingPage&groups=C0002%3A1%2CC0009%3A1%2CC0007%3A1%2CC0003%3A1%2CC0001%3A1%2CC0004%3A1&hosts=H138%3A1%2CH29%3A1%2CH111%3A1%2CH129%3A1%2CH93%3A1%2CH128%3A1%2CH112%3A1%2CH147%3A1%2CH148%3A1%2CH56%3A1%2CH58%3A1%2CH59%3A1%2CH91%3A1%2CH20%3A1%2CH104%3A1%2CH115%3A1%2CH75%3A1%2CH86%3A1%2CH25%3A1%2CH90%3A1%2CH32%3A1%2CH116%3A1%2CH124%3A1%2CH7%3A1%2CH152%3A1%2CH37%3A1%2CH42%3A1%2CH43%3A1%2CH153%3A1%2CH149%3A1%2CH145%3A1%2CH134%3A1%2CH139%3A1%2CH144%3A1&genVendors=V77%3A1%2CV67%3A1%2CV79%3A1%2CV71%3A1%2CV69%3A1%2CV7%3A1%2CV5%3A1%2CV9%3A1%2CV1%3A1%2CV70%3A1%2CV3%3A1%2CV68%3A1%2CV78%3A1%2CV17%3A1%2CV76%3A1%2CV80%3A1%2CV16%3A1%2CV72%3A1%2CV10%3A1%2CV40%3A1%2C&geolocation=PK%3BPB&AwaitingReconsent=false; _uetsid=7c2e4db0213511f0bca4851cb46f5236|1i11fu8|2|fve|0|1940; ttcsid=1745665466430::vrjZY0NGJk-CRjfo3x5_.8.1745668362756; _scid_r=r4pt02EUMDwDAF4F9vfx7LsfKVHQYORBse8OfA; _sc_cspv=https%3A%2F%2Ftr.snapchat.com%2Fconfig%2Fcom%2F5df43118-abd2-4cd8-89b9-8cf942d1ee25.js%3Fv%3D3.44.6-2504241707; _ga_9J2BFGDX1E=GS1.1.1745665465.10.1.1745668362.58.0.0; _ga=GA1.1.2084360745.1745402880; ttcsid_CJ5M5PJC77U7DSNBELOG=1745665466430::YRashA3ej7oGtyT472uB.8.1745668362967; _uetvid=d67945d0202a11f0817e972d9bec397e|1b5wek3|1745668363515|12|1|bat.bing.com/p/insights/c/d")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(bodyText, &jsonResponse); err != nil {
		log.Fatal(err)
	}
	return jsonResponse
}
func main() {
	client := &http.Client{}
	start := 1
	end := 9999999
	file, err := os.Create("data.json")
	if err != nil {
		log.Fatal("Failed to create JSON file:", err)
	}
	defer file.Close()
	file.WriteString("[\n")
	first := true
	for wc := start; wc <= end; wc++ {
		url := fmt.Sprintf("https://apigw.trendyol.com/discovery-sfint-browsing-service/api/search-feed/products?source=sr?wc=%d&size=60", wc)
		fmt.Println("Fetching products for wc:", wc)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("accept", "application/json")
		req.Header.Set("accept-language", "en-US,en;q=0.9")
		req.Header.Set("baggage", "ty.kbt.name=ViewSearchResult,ty.platform=Web,ty.business_unit=Core%20Commerce,ty.channel=INT,com.trendyol.observability.business_transaction.name=ViewSearchResult,ty.source.service.name=WEB%20Storefront%20International,ty.source.deployment.environment=production,ty.source.service.version=4f92d141,ty.source.client.path=unknown,ty.source.service.type=client")
		req.Header.Set("content-type", "application/json")
		req.Header.Set("origin", "https://www.trendyol.com")
		req.Header.Set("platform", "Web")
		req.Header.Set("priority", "u=1, i")
		req.Header.Set("sec-ch-ua", `"Google Chrome";v="135", "Not-A.Brand";v="8", "Chromium";v="135"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"macOS"`)
		req.Header.Set("sec-fetch-dest", "empty")
		req.Header.Set("sec-fetch-mode", "cors")
		req.Header.Set("sec-fetch-site", "same-site")
		req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/135.0.0.0 Safari/537.36")
		req.Header.Set("x-section-id", "null")
		req.Header.Set("cookie", "csrf_secret=u2ANLHbZzCsSCZr5y4L3ziOG; platform=web; _cfuvid=JS6Mjd5imwMhUE.DEVCgMHGZxxyaowoBsHkO6AJXN.w-1745402857084-0.0.1.1-604800000; anonUserId=c88ca0a0-202a-11f0-b282-e524755a5054; __cflb=04dToYCH9RsdhPpttDDEnPngTWcVjd8n2VS1BNR3Xj; OptanonAlertBoxClosed=2025-04-23T10:07:40.439Z; pid=c88ca0a0-202a-11f0-b282-e524755a5054; sid=RN4GQHDizD; storefrontId=36; countryCode=AE; language=en; functionalConsent=true; performanceConsent=true; targetingConsent=true; WebRecoTss=similarRecoAdsVersion%2F1%7CallInOneRecoVersion%2F1%7CbasketRecoVersion%2F1%7CshopTheLookVersion%2F1%7CpdpGatewayVersion%2F1%7CsimilarSameBrandVersion%2F1%7CcrossSameBrandVersion%2F1%7CcrossRecoVersion%2F1%7CsimilarRecoVersion%2F1%7CcompleteTheLookVersion%2F1%7CcollectionRecoVersion%2F1%7CcrossRecoAdsVersion%2F1%7CnavigationSideMenuVersion%2FsideMenu%3AinitialTest_1(M)%7ChomepageVersion%2Fsorter%3AhomepageSorterNewTest_d.firstComponent%3AinitialNewTest_1(M)%7CnavigationSectionVersion%2Fsection%3AazSectionTest_1(M)%7CfirstComponent_V1%2F1%7Csorter_V1%2Fb%7Csection_V1%2F1%7CsideMenu_V1%2F1%7CtopWidgets_V1%2F1; UserInfo=%7B%22Gender%22%3Anull%2C%22UserTypeStatus%22%3Anull%2C%22ForceSet%22%3Afalse%7D; navbarGenderId=1; _gcl_au=1.1.140209073.1745402880; tss=firstComponent_V1_1%2Csorter_V1_b%2Csection_V1_1%2CsideMenu_V1_1%2CtopWidgets_V1_1%2CFSA_B%2CProductCardVariantCount_B%2CSuggestionPopular_B%2CRR_2%2CGRRLO_B%2CGRRIn_B%2CVisualCategorySlider_B%2CSuggestionTermActive_B%2CKB_B%2CDGB_B%2CSB_B%2CSuggestion_B%2COFIR_B; _gid=GA1.2.1617466369.1745402880; _scid=kApt02EUMDwDAF4F9vfx7LsfKVHQYORB; _fbp=fb.1.1745402880754.810322442114590411; _ScCbts=%5B%5D; _pin_unauth=dWlkPU5UZ3hNREptTWpJdE1qZ3hPUzAwT0RFMExXSXlOR010TlRreFpESXdOMlZoWTJVNA; _tt_enable_cookie=1; _ttp=01JSH1WRWS90XBGHC34CSJZ9SE_.tt.1; AbTestingCookies=A_82-B_38-C_43-D_9-E_38-F_99-G_63-H_9-I_13-J_94-K_50-L_55-M_17-N_13-O_80; hvtb=1; VisitCount=1; SearchMode=1; WebAbTesting=A_16-B_93-C_88-D_50-E_32-F_13-G_69-H_13-I_35-J_51-K_50-L_84-M_21-N_65-O_66-P_10-Q_76-R_88-S_88-T_73-U_72-V_88-W_6-X_53-Y_65-Z_31; ForceUpdateSearchAbDecider=forced; WebRecoAbDecider=ABbasketRecoVersion_1-ABcrossRecoVersion_1-ABcrossRecoAdsVersion_1-ABsimilarRecoVersion_1-ABsimilarSameBrandVersion_1-ABcrossSameBrandVersion_1-ABpdpGatewayVersion_1-ABallInOneRecoVersion_1-ABattributeRecoVersion_1-ABcollectionRecoVersion_1-ABshopTheLookVersion_1-ABsimilarRecoAdsVersion_1-ABcompleteTheLookVersion_1-ABhomepageVersion_firstComponent%3AinitialNewTest_1.performanceSorting%3AinitialTest_3.sorter%3AhomepageSorterNewTest_d%28M%29-ABnavigationSideMenuVersion_sideMenu%3AinitialTest_1%28M%29-ABnavigationSectionVersion_section%3AazSectionTest_1%28M%29; FirstSession=0; msearchAb=ABAdvertSlotPeriod_1-ABDsNlp_2-ABQR_B-ABSearchFETestV1_B-ABBSA_D-ABSuggestionLC_B; WebAbDecider=ABres_B-ABBMSA_B-ABRRIn_B-ABSCB_B-ABSuggestionHighlight_B-ABBP_A-ABCatTR_B-ABSuggestionTermActive_A-ABAZSmartlisting_63-ABBH2_B-ABMB_B-ABMRF_1-ABARR_B-ABMA_B-ABSP_B-ABPastSearches_B-ABSuggestionJFYProducts_B-ABSuggestionQF_B-ABBadgeBoost_A-ABRelevancy_1-ABFilterRelevancy_1-ABSuggestionBadges_B-ABProductGroupTopPerformer_B-ABOpenFilterToggle_2-ABRR_2-ABBS_2-ABSuggestionPopularCTR_B; AbTesting=pdpAiReviewSummaryUat_B-SFWBAA_V1_B-SFWDBSR_A-SFWDQL_B-SFWDRS_A-SFWDSAOFv2_B-SFWDSFAG_B-SFWDTKV2_A-SSTPRFL_B-STSBuynow_B-STSCouponV2_A-STSImageSocialProof_A-STSRecoRR_B-STSRecoSocialProof_A-WCBsQckFiltTestv2_B-WCOnePageCheckout_B-WEBSFAATest1_A-WebSFAATest2_B-WebSFAATest3_A%7C1745405135%7Cc88ca0a0-202a-11f0-b282-e524755a5054; __cf_bm=NWXTz_0PnBOy.pCO9V30IFpkTw2Qo2Rje1YU4o55T9k-1745403356-1.0.1.1-iyIpH5eFby7QY4AOTWZym26XHjefMDMuC379iE7HNBCR9f5EbPoplMYfMFa5Me.CvsRrbR079unZJVptaKiMm4qrDsLxMA1Jbq3tngIegkc; _ga_9J2BFGDX1E=GS1.1.1745402880.1.1.1745403367.7.0.0; _scid_r=qopt02EUMDwDAF4F9vfx7LsfKVHQYORBse8OLA; _ga=GA1.2.2084360745.1745402880; _dc_gtm_UA-13174585-70=1; _sc_cspv=https%3A%2F%2Ftr.snapchat.com%2Fconfig%2Fcom%2F5df43118-abd2-4cd8-89b9-8cf942d1ee25.js%3Fv%3D3.44.3-2504222057; OptanonConsent=isGpcEnabled=0&datestamp=Wed+Apr+23+2025+15%3A16%3A07+GMT%2B0500+(Pakistan+Standard+Time)&version=202402.1.0&browserGpcFlag=0&isIABGlobal=false&consentId=92cde421-80dd-4b6c-9752-f4bc74338ad3&interactionCount=1&isAnonUser=1&landingPath=NotLandingPage&groups=C0002%3A1%2CC0009%3A1%2CC0007%3A1%2CC0003%3A1%2CC0001%3A1%2CC0004%3A1&hosts=H138%3A1%2CH29%3A1%2CH111%3A1%2CH129%3A1%2CH93%3A1%2CH128%3A1%2CH112%3A1%2CH147%3A1%2CH148%3A1%2CH56%3A1%2CH58%3A1%2CH59%3A1%2CH91%3A1%2CH20%3A1%2CH104%3A1%2CH115%3A1%2CH75%3A1%2CH86%3A1%2CH25%3A1%2CH90%3A1%2CH32%3A1%2CH116%3A1%2CH124%3A1%2CH7%3A1%2CH152%3A1%2CH37%3A1%2CH42%3A1%2CH43%3A1%2CH153%3A1%2CH149%3A1%2CH145%3A1%2CH134%3A1%2CH139%3A1%2CH144%3A1&genVendors=V77%3A1%2CV67%3A1%2CV79%3A1%2CV71%3A1%2CV69%3A1%2CV7%3A1%2CV5%3A1%2CV9%3A1%2CV1%3A1%2CV70%3A1%2CV3%3A1%2CV68%3A1%2CV78%3A1%2CV17%3A1%2CV76%3A1%2CV80%3A1%2CV16%3A1%2CV72%3A1%2CV10%3A1%2CV40%3A1%2C&geolocation=PK%3BPB&AwaitingReconsent=false; _uetsid=d6791910202a11f097b9bfaf0c52dfc1|n1586e|2|fvb|0|1939; cto_bundle=Bag8ol9DNGtMb3BmVDRnNWVnNWJkMTVQcmFDb0olMkYlMkIwNmZuZGpqR2JqS2VmamxjZWltJTJCWEUxdmw3cXJwTWhEeGFzelBNdU83dXFjSG1CemZpOFJrVkpFJTJGelMyNEk0SkhWZG9DcyUyQnZmYlpmM0JBd1JLMDg2TmIzYzR3MDJMMDdDczJaaHRnQnpoWCUyRm1ib1RHTGV1QzZVTVlqOGNOQm51THYwSEdmeUI5aVZENDdDeVhuRkslMkJtUkJwV2Q5cm1NYWlyQWFwaTFrY3c2cmtZWiUyRm8xN05QNVZEZiUyRmRBJTNEJTNE; ttcsid=1745402880923::pm5yFa6t5a51-RWmpBHn.1.1745403368456; ttcsid_CJ5M5PJC77U7DSNBELOG=1745402880923::I82TPjx5GfCIEgRdpsw1.1.1745403368663; _uetvid=d67945d0202a11f0817e972d9bec397e|1jcuqwt|1745403369059|3|1|bat.bing.com/p/insights/c/u")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result Root
		if err := json.Unmarshal(bodyText, &result); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Total products found:", len(result.Data.Contents))
		if len(result.Data.Contents) == 0 {
			fmt.Println("No more products found ")
			continue
		}
		for _, p := range result.Data.Contents {
			time.Sleep(4 * time.Second)
			fmt.Println("Fetching details for product ID:", p.ID)
			detailedProduct := fetchProductDetails(p.ID)
			encoder := json.NewEncoder(file)
			if !first {
				file.WriteString(",\n")
			}
			first = false
			if err := encoder.Encode(detailedProduct); err != nil {
				log.Printf("Failed to write product ID %d to file: %v", p.ID, err)
				continue
			}
		}
	}
	file.WriteString("\n]")
	fmt.Println("Done!")
}
