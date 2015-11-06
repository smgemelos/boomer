# Database Backend
**Version 1.0**

## User Table Api:
    AddUser:  
        Type:   POST
        Path:   /user
        Data:   {"level":int, "username":"string", "password":"sting"}

        Example:    curl -X POST -d '{"level":2,"username":"Steven","password":"stevenpass"}' http://localhost:3000/user
        Response:   {"userid":1,"created":"2015-08-05 11:31:59","updated":"2015-08-05 11:31:59","level":2,"username":"Steven","password":"stevenpass"}

    GetUser: 
        Type:   GET
        Path:   /user/{Userid}
        Data:   

        Example:    curl -X GET http://localhost:3000/user/1
        Response:   {"userid":1,"created":"2015-08-05 11:31:59","updated":"2015-08-05 11:31:59","level":2,"username":"Steven","password":"stevenpass"}

	GetUsers:
        Type:   GET
        Path:   /user
        Data:   where filters include any of {"level":int, "username":"string", "password":"sting"}

        Example:    curl -X GET -d '{"level":2}' http://localhost:3000/user
        Response:   [{"userid":1,"created":"2015-08-05 11:31:59","updated":"2015-08-05 11:31:59","level":2,"username":"Steven","password":"stevenpass"},
                     {"userid":3,"created":"2015-08-05 11:31:59","updated":"2015-08-05 11:31:59","level":2,"username":"Patrick","password":"patrickpass"}]

	UpdateUser:
        Type:   POST
        Path:   /user/{Userid}
        Data:   {"userid":int, level":int, "username":"string", "password":"sting"}

        Example:    curl -X POST -d '{"userid":1,"level":2,"username":"Steven","password":"stevenpass"}' http://localhost:3000/user
        Response:   {"userid":1,"created":"2015-08-05 11:31:59","updated":"2015-08-05 11:31:59","level":2,"username":"Steven","password":"stevenpass"}


	DeleteUser:
        Type:   DELETE
        Path:   /user/{Userid}
        Data:   

        Example:    curl -X DELETE http://localhost:3000/user/1
        Response:   


	GetUserCpe:
        Type:   GET
        Path:   /user/{Userid}/cpe
        Data:   

        Example:    curl -X GET http://localhost:3000/user/5/cpe
        Response:   [{"cpeid":1,"created":"2015-08-05 11:31:59","updated":"2015-08-05 11:31:59","new":true,"userid":5,
                      "topoid":0,"ipaddr":"192.168.10.10","macaddr":"32:20:39:26:00:00","publicip":"10.250.3.62","name":"","location":""},
                     {"cpeid":2,"created":"2015-08-05 11:31:59","updated":"2015-08-05 11:31:59","new":false,"userid":5,
                      "topoid":0,"ipaddr":"192.168.10.11","macaddr":"32:20:39:26:00:01","publicip":"10.250.3.64","name":"","location":""}]

## CPE Table Api:
    AddCpe:  
        Type:   POST
        Path:   /cpe
        Data:   {"new":bool, "userid":int, "topoid":int, "macaddr":string, "mgmtip":string, "mgmtgw":string, "publicip":string, "publicgw":string, "name":string, "location":string}

        Example:    curl -X POST -d '{"new":true,"userid":5,"topoid":0,"macaddr":"32:20:39:26:00:00","mgmtip":"192.168.10.10/24","mgmtgw":"192.168.10.1","publicip":"10.250.3.62/16","publicgw":"10.250.0.1","name":"steveCPE1","location":"stevehome"}' http://localhost:3000/cpe 

        Response:   {"cpeid":1,"created":"2015-08-05 14:06:49","updated":"2015-08-05 14:06:49","new":true,"userid":5,"topoid":0,"ipaddr":"192.168.10.10/24",
                     "macaddr":"32:20:39:26:00:00","publicip":"10.250.3.62/16","publicgw":"10.250.0.1","name":"steveCPE1","location":"stevehome"}

    GetCpe: 
        Type:   GET
        Path:   /cpe/{Cpeid}
        Data:   

        Example:    curl -X GET http://localhost:3000/cpe/1

        Response:   {"cpeid":1,"created":"2015-08-05 14:06:49","updated":"2015-08-05 14:06:49","new":true,"userid":5,"topoid":0,
                     "macaddr":"32:20:39:26:00:00","mgmtip":"192.168.10.10/24","mgmtgw":"192.168.10.1","publicip":"10.250.3.62/16","publicgw":"10.250.0.1","name":"steveCPE1","location":"stevehome"}

    GetCpes:
        Type:   GET
        Path:   /cpe
        Data:   Where fitlers include any of {"new":bool, "userid":int, "topoid":int, "ipaddr":string, "macaddr":string, "publicip":string, "publicgw":string, "name":string, "location":string}

        Example:    curl -X GET -d '{"userid":5}' http://localhost:3000/cpe
        Response:   [{"cpeid":1,"created":"2015-08-05 14:06:49","updated":"2015-08-05 14:06:49","new":true,"userid":5,"topoid":0,"macaddr":"32:20:39:26:00:00","mgmtip":"192.168.10.10/24","mgmtgw":"192.168.10.1","publicip":"10.250.3.62/16","publicgw":"10.250.0.1","name":"","location":""},
                     {"cpeid":2,"created":"2015-08-05 14:06:49","updated":"2015-08-05 14:06:49","new":false,"userid":5,"topoid":0,
                      "macaddr":"32:20:39:26:00:01","mgmtip":"192.168.10.11/24","mgmtgw":"192.168.10.1","publicip":"10.250.3.64.16","publicgw":"10.250.0.1","name":"","location":""}]

        Note:   Gets all CPEs based on where filters

    UpdateCpe:
        Type:   POST
        Path:   /cpe/{Cpeid}
        Data:   {"userid":int, level":int, "username":"string", "password":"sting"}

        Example:    curl -X POST -d '{"cpeid":3,"new":false,"userid":1,"topoid":0,"ipaddr":"192.168.10.15","macaddr":"32:20:39:26:00:02","publicip":"10.250.3.62","name":"MyCpe","location":"office"}' http://localhost:3000/cpe/3

        Response:    {"cpeid":3,"created":"2015-08-05 14:06:49","updated":"2015-08-05 14:06:49","new":false,"userid":1,"topoid":0,
                      "macaddr":"32:20:39:26:00:01","mgmtip":"192.168.10.10/24","mgmtgw":"192.168.10.1","publicip":"10.250.3.64/16","publicgw":"10.250.0.1","name":"MyCpe","location":"Office"}


    DeleteCpe:
        Type:   DELETE
        Path:   /cpe/{Cpeid}
        Data:   

        Example:    curl -X DELETE http://localhost:3000/cpe/1
        Response:   


    GetCpeTopo:
        Type:   GET
        Path:   /cpe/{Cpeid}/topo
        Data:   

        Example:    curl -X GET http://localhost:3000/cpe/5/topo
        Response:   {"topoid":3,"created":"2015-08-05 14:30:28","updated":"2015-08-05 14:30:28","cpeid":5,"userid":5,"name":"basictopo","topo":JSON_String}


## Topo Table Api:
    AddTopo:  
        Type:   POST
        Path:   /topo
        Data:   {"cpeid":int64, "userid":int64, "name":string, "topogui":JSON_String, "topoorch":JSON_String}

        Example:    curl -X POST -d '{"cpeid":1,"userid":5,"name":"basictopo","topogui":"{JSON}","topoorch":"{JSON}"}' http://localhost:3000/topo

        Response:   {"topoid":1,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","cpeid":1,"userid":5,"name":"basictopo","topogui":"{JSON}","topoorch":"{JSON}"}

        Note:   Updates CPE Table entry for indicated cpeid to map to this topoid

    GetTopo: 
        Type:   GET
        Path:   /topo/{Topoid}
        Data:   

        Example:    curl -X GET http://localhost:3000/topo/1

        Response:   {"topoid":1,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","cpeid":1,"userid":5,"name":"basictopo","topogui":"{JSON}","topoorch":"{JSON}"}

    GetTopos: 
        Type:   GET
        Path:   /topo
        Data:   Where Filters {"cpeid":int64, "userid":int64, "name":string, "topogui":JSON_String, "topoorch":JSON_String}

        Example:    curl -X GET http://localhost:3000/topo
        Response:   [{"topoid":1,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","cpeid":1,"userid":5,"name":"basictopo","topogui":"topology","topoorch":"topology"},
                     {"topoid":2,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","cpeid":2,"userid":5,"name":"basictopo2","topogui":"topology2","topoorch":"topology2"},
                     {"topoid":3,"created":"2015-08-05 14:39:28","updated":"2015-08-05 14:39:28","cpeid":5,"userid":5,"name":"basictopo3","topogui":"topology3","topoorch":"topology3"},
                     {"topoid":4,"created":"2015-08-05 14:39:28","updated":"2015-08-05 14:39:28","cpeid":4,"userid":5,"name":"basictopo4","topogui":"topology4","topoorch":"topology4"}]

        Note:   Gets all the topos mathcing Where filters

    UpdateTopo:
        Type:   POST
        Path:   /topo/{Topoid}
        Data:   {"topoid":1,"cpeid":1,"userid":5,"name":"basictopo","topo":"{JSON}"}

        Example:    curl -X POST -d '{"topoid":4,"cpeid":4,"userid":5,"name":"advancedtopo1","topo":"topology4"}' http://localhost:3000/topo/4

        Response:   {"topoid":4,"created":"2015-08-05 14:39:28","updated":"2015-08-05 14:39:28","cpeid":4,"userid":5,"name":"advancedtopo1","topogui":"{topology4}","topoorch":"{topology4}"}


    DeleteTopo:
        Type:   DELETE
        Path:   /topo/{Topoid}
        Data:   

        Example:    curl -X DELETE http://localhost:3000/topo/1
        Response:   


## Vnf Table Api:
    AddVnf:  
        Type:   POST
        Path:   /vnf
        Data:   {"name":string, "vnftype":string, "favor":string, "image":string, "version":string, "template":string, "distro":string, "release":string, "arch":string, "startcmd":string}

        Example:    curl -X POST -d '{"name":"Linux Router","vnftype":"router","flavor":"docker","image":"ubuntu14.04","version":"trusty",
                                      "template":"", "distro":"", "release":"", "arch":"", "startcmd":"/bin/bash"}' http://localhost:3000/vnf

        Response:   {"vnfid":1,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","name":"Linux Router","vnftype":"router",
                     "flavor":"docker","image":"ubuntu","version":"latest","template":"","distro":"","release":"","arch":"","startsmd":"/bin/bash"}

    GetVnf: 
        Type:   GET
        Path:   /vnf/{Vnfid}
        Data:   

        Example:    curl -X GET http://localhost:3000/topo/1

        Response:   {"vnfid":1,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","name":"Linux Router","vnftype":"router","contype":"docker","image":"ubuntu14.04","version":"2.0","startfile":"/bin/bash"}

    GetVnfs: 
        Type:   GET
        Path:   /vnf
        Data:   Where filters  {"name":string, "vnftype":string, "favor":string, "image":string, "version":string, "template":string, "distro":string, "release":string, "arch":string, "startcmd":string}  

        Example:    curl -X GET http://localhost:3000/topo
        Response:   [{"vnfid":1,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","name":"Linux Router","vnftype":"router","contype":"flavor",
                      "image":"ubuntu","version":"latest","template":"","distro":"","release":"","arch":"","startsmd":"/bin/bash"},
                     {"vnfid":2,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","name":"Linux Firewall","vnftype":"firewall","contype":"docker",
                      "image":"ubuntu","version":"latest","template":"","distro":"","release":"","arch":"","startsmd":"/bin/bash"},
                     {"vnfid":3,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","name":"Appache Traffic Server","vnftype":"cache","contype":"docker",
                      "image":"ubuntu","version":"latest","template":"","distro":"","release":"","arch":"","startsmd":""},
                     {"vnfid":4,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","name":"Garbage","vnftype":"cache","contype":"docker",
                      "image":"ubuntu","version":"latest","template":"","distro":"","release":"","arch":"","startsmd":"/bin/bash"}]

        Note:   Gets all the vnfs mathcinh Where filters

    UpdateVnf:
        Type:   POST
        Path:   /vnf/{Vnfid}
        Data:   {"vnfid":4,"name":"Garbage","vnftype":"cache","contype":"docker","image":"ubuntu14.04","version":"2.0","startfile":""}

        Example:    curl -X POST -d '{"vnfid":4,"name":"Garbage2","vnftype":"cache","contype":"docker","image":"ubuntu14.04","version":"2.0","startfile":""}' http://localhost:3000/vnf/4

        Response:   {"vnfid":4,"created":"2015-08-05 14:39:27","updated":"2015-08-05 14:39:27","name":"Garbage2","vnftype":"cache","contype":"docker","image":"ubuntu14.04","version":"2.0","startfile":""}


    DeleteVnf:
        Type:   DELETE
        Path:   /vnf/{Vnfid}
        Data:   

        Example:    curl -X DELETE http://localhost:3000/vnf/1
        Response:   






