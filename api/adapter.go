package api

func ModifyServicesToSliceOfStringSlices(services []Service, ipAndPortSearchMode string, columnsHeaders []string) [][]string {
	if len(services) == 0 {
		return nil
	}

	//TODO: remove all append functions
	preparedData := [][]string{}
	if ipAndPortSearchMode != "nope" {
		// TODO: refator that, below same code
		service := findServiceForIPAndPortSearchMode(services, ipAndPortSearchMode)
		vip := service.ServiceIP + ":" + service.ServicePort
		serviceState := ""
		if service.IsUp {
			serviceState = "UP"
		} else {
			serviceState = "DOWN"
		}
		balanceType := service.BalanceType
		routingType := ""
		switch service.RoutingType {
		case "masquerading":
			routingType = "nat"
		case "tunneling":
			routingType = "tun"
		}
		timeout := service.Healthcheck.Timeout.String()
		repeatHealthcheck := service.Healthcheck.RepeatHealthcheck.String()
		tAndR := timeout + "/" + repeatHealthcheck
		protocol := service.Protocol

		for _, applicationServer := range service.ApplicationServers {
			appSrvIPPort := applicationServer.ServerIP + ":" + applicationServer.ServerPort
			serverState := ""
			if applicationServer.IsUp {
				serverState = "UP"
			} else {
				serverState = "DOWN"
			}
			srvHCType := service.Healthcheck.Type
			srvHCAddr := applicationServer.ServerHealthcheck.HealthcheckAddress
			data := []string{vip, serviceState, appSrvIPPort, serverState, protocol, routingType, balanceType, srvHCType, srvHCAddr, tAndR}
			preparedData = append(preparedData, data)
		}
		return preparedData
	}

	for i, service := range services {
		vip := service.ServiceIP + ":" + service.ServicePort
		serviceState := ""
		if service.IsUp {
			serviceState = "UP"
		} else {
			serviceState = "DOWN"
		}
		balanceType := service.BalanceType
		routingType := ""
		switch service.RoutingType {
		case "masquerading":
			routingType = "nat"
		case "tunneling":
			routingType = "tun"
		}
		timeout := service.Healthcheck.Timeout.String()
		repeatHealthcheck := service.Healthcheck.RepeatHealthcheck.String()
		tAndR := timeout + "/" + repeatHealthcheck
		protocol := service.Protocol
		for _, applicationServer := range service.ApplicationServers {
			appSrvIPPort := applicationServer.ServerIP + ":" + applicationServer.ServerPort
			serverState := ""
			if applicationServer.IsUp {
				serverState = "UP"
			} else {
				serverState = "DOWN"
			}
			srvHCType := service.Healthcheck.Type
			srvHCAddr := applicationServer.ServerHealthcheck.HealthcheckAddress
			data := []string{vip, serviceState, appSrvIPPort, serverState, protocol, routingType, balanceType, srvHCType, srvHCAddr, tAndR}
			preparedData = append(preparedData, data)
		}
		if i != len(services)-1 {
			emptyTableData := columnsHeaders
			preparedData = append(preparedData, emptyTableData)
		}
	}
	return preparedData
}

func findServiceForIPAndPortSearchMode(services []Service, ipAndPortSearchMode string) *Service {
	for _, service := range services {
		if service.ServiceIP+":"+service.ServicePort == ipAndPortSearchMode {
			return &service
		}
		for _, applicationServer := range service.ApplicationServers {
			if applicationServer.ServerIP+":"+applicationServer.ServerPort == ipAndPortSearchMode {
				return &service
			}
		}
	}
	return nil
}
