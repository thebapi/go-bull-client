package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
    GENERATE_CUSTOM_ASSET_SUMMARY_REPORT: 'GENERATE_CUSTOM_ASSET_SUMMARY_REPORT',
    GENERATE_CUSTOM_ASSET_JOURNEY_REPORT: 'GENERATE_CUSTOM_ASSET_JOURNEY_REPORT',
    GENERATE_CUSTOM_ASSET_PERFORMANCE_REPORT: 'GENERATE_CUSTOM_ASSET_PERFORMANCE_REPORT',
    GENERATE_CUSTOM_ASSET_INCIDENT_REPORT: 'GENERATE_CUSTOM_ASSET_INCIDENT_REPORT',
    GENERATE_CUSTOM_SITE_PERFORMANCE_REPORT: 'GENERATE_CUSTOM_SITE_PERFORMANCE_REPORT',
    GENERATE_CUSTOM_ZONE_PERFORMANCE_REPORT: 'GENERATE_CUSTOM_ZONE_PERFORMANCE_REPORT',
    GENERATE_CUSTOM_COMPANY_PERFORMANCE_REPORT: 'GENERATE_CUSTOM_COMPANY_PERFORMANCE_REPORT',

 */

func main() {
	addAssetSummaryReportJobs();
	addAssetJourneyReportJobs();
	addAssetIncidentReportJobs();
	addAssetPerformanceReportJobs();
	addCustomSiteReportJob();
	addCustomZoneReportJob();
	addCustomCompanyReportJob();
}


func addAssetSummaryReportJobs (){
	jsonData := map[string]string{
		"query": `
            mutation {
				  addCustomAssetReportJob(assetIds: [214], from: 1601664493, to: 1606502893,createdBy: 99, reportType: "GENERATE_CUSTOM_ASSET_SUMMARY_REPORT") {
					success
					message
					jobIds
				  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	postRequest(jsonValue)
}

func addAssetJourneyReportJobs (){
	jsonData := map[string]string{
		"query": `
            mutation {
				  addCustomAssetReportJob(assetIds: [214], from: 1601664493, to: 1606502893,createdBy: 99, reportType: "GENERATE_CUSTOM_ASSET_JOURNEY_REPORT") {
					success
					message
					jobIds
				  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	postRequest(jsonValue)
}

func addAssetIncidentReportJobs (){
	jsonData := map[string]string{
		"query": `
            mutation {
				  addCustomAssetReportJob(assetIds: [214], from: 1601664493, to: 1606502893,createdBy: 99, reportType: "GENERATE_CUSTOM_ASSET_INCIDENT_REPORT") {
					success
					message
					jobIds
				  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	postRequest(jsonValue)
}

func addAssetPerformanceReportJobs (){
	jsonData := map[string]string{
		"query": `
            mutation {
				  addCustomAssetReportJob(assetIds: [214], from: 1601664493, to: 1606502893,createdBy: 99, reportType: "GENERATE_CUSTOM_ASSET_PERFORMANCE_REPORT") {
					success
					message
					jobIds
				  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	postRequest(jsonValue)
}

func addCustomSiteReportJob (){
	jsonData := map[string]string{
		"query": `
            mutation {
				  addCustomSiteReportJob(siteIds: [2], from: 1601664493, to: 1606502893,createdBy: 99, reportType: "GENERATE_CUSTOM_SITE_PERFORMANCE_REPORT") {
					success
					message
					jobIds
				  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	postRequest(jsonValue)
}

func addCustomZoneReportJob (){
	jsonData := map[string]string{
		"query": `
            mutation {
				  addCustomZoneReportJob(zoneIds: [8], from: 1601664493, to: 1606502893,createdBy: 99, reportType: "GENERATE_CUSTOM_ZONE_PERFORMANCE_REPORT") {
					success
					message
					jobIds
				  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	postRequest(jsonValue)
}


func addCustomCompanyReportJob (){
	jsonData := map[string]string{
		"query": `
            mutation {
				  addCustomCompanyReportJob(companyIds: [2], from: 1601664493, to: 1606502893,createdBy: 99, reportType: "GENERATE_CUSTOM_COMPANY_PERFORMANCE_REPORT") {
					success
					message
					jobIds
				  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	postRequest(jsonValue)
}



func postRequest(jsonValue []byte) {
	resp, err := http.Post("http://0.0.0.0:4848/", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func addNotificationJobs (){
	jsonData := map[string]string{
		"query": `
            mutation {
				  addIncidentNotificationJob(ids: [36384, 36368, 36366]) {
					success
					jobIds
				  }
			}
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)

	resp, err := http.Post("http://0.0.0.0:4848/", "application/json",  bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}