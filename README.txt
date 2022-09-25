---- Awair assignment ----

@Requirements

Given a file with devices data (devices.json), write 2 services Client API and DeviceService

Client API should parse the JSON file and have REST interface

The file is of unknown size, it can contain several millions of records. Client API has limited resources available (e.g. 200MB ram)

While reading the file, it should call a second service (DeviceService), that either creates a new record in a database or updates the existing one

Client API should provide an endpoint to retrieve the data from the second service (DeviceService)

Get a device by ID

Get devices by type (with pagination)

Get connected or disconnected devices (with pagination)

List All devices (with pagination)

The end result should be a database containing the devices, representing the latest version found in the JSON. (Database can be memory)

@ Planning 

1- Database connection - In progress

2- Parsing all devices data from devices json - In progress

3- Implementing DeviceService 

	+ Get a device by ID

	+ Get devices by type (with pagination)

	+ Get connected or disconnected devices (with pagination)

	+ List All devices (with pagination)

	+ Insert/update in bulk

@ Notes:

	+ Since it is a long time that I do not touch Go so I may need more time to complete this task.

	+ 2022-09-26: 2 hours spended
		- Add deviceInfoTransform func to transform raw data to Device struct
		- Add insertDevices to insert new devices in batch (currently notworking)     
