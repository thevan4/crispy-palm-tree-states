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
		vip := service.IP + ":" + service.Port
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
		timeout := service.HCTimeout.String()
		repeatHealthcheck := service.HCRepeat.String()
		tAndR := timeout + "/" + repeatHealthcheck
		protocol := service.Protocol

		for _, applicationServer := range service.ApplicationServers {
			appSrvIPPort := applicationServer.IP + ":" + applicationServer.Port
			serverState := ""
			if applicationServer.IsUp {
				serverState = "UP"
			} else {
				serverState = "DOWN"
			}
			srvHCType := service.HCType
			srvHCAddr := applicationServer.HCAddress
			data := []string{vip, serviceState, appSrvIPPort, serverState, protocol, routingType, balanceType, srvHCType, srvHCAddr, tAndR}
			preparedData = append(preparedData, data)
		}
		return preparedData
	}

	for i, service := range services {
		vip := service.IP + ":" + service.Port
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
		timeout := service.HCTimeout.String()
		repeatHealthcheck := service.HCRepeat.String()
		tAndR := timeout + "/" + repeatHealthcheck
		protocol := service.Protocol
		for _, applicationServer := range service.ApplicationServers {
			appSrvIPPort := applicationServer.IP + ":" + applicationServer.Port
			serverState := ""
			if applicationServer.IsUp {
				serverState = "UP"
			} else {
				serverState = "DOWN"
			}
			srvHCType := service.HCType
			srvHCAddr := applicationServer.HCAddress
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
		if service.IP+":"+service.Port == ipAndPortSearchMode {
			return &service
		}
		for _, applicationServer := range service.ApplicationServers {
			if applicationServer.IP+":"+applicationServer.Port == ipAndPortSearchMode {
				return &service
			}
		}
	}
	return nil
}
