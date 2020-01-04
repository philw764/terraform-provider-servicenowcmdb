Terraform Provider for ServiceNow
==================

A Terraform Provider to provide the ability to add and/or update Configuration Items (CIs) in the ServiceNow Configuration Management Database (CMDB).  This provider enables you to align 
Infrastructure built using Terraform in the Cloud or on Premise and the ServiceNow CMDB. Alignment of your Infrastructure and ServiceNow CMDB provides greater levels of control using Configuration Management disicplines.

Why is this important?  ServiceNow already has the ability to discover infrastructure and configuration items from AWS and Azure now so why would we need this provider?  Having the ability to perform CRUD operations on Configuration Items
in the ServiceNow CMDB that correspond to your Infrastructure as Code (IaC) activities provides you with the ability ensure that the CMDB accurately reflects the infrastructure being defined plus it enables 
you to add additional information such as relationships, reference linking and custom field values that are just not able to be captured using automatic discovery tools.

<img src="https://www.terraform.io/assets/images/og-image-8b3e4f7d.png" height="200" alt="Terraform Logo"/><img src="https://community.servicenow.com/c4fe846adbb95f0037015e77dc961918.iix" height="200" alt="ServiceNow Logo"/>

Quick Overview
==============
The very simple example below shows a Terraform Script (HCL Script) with the ServiceNow Provider **servicenowcmdb** and a single resource called **testci4**.  This script will create, modify or destroy
a Configuration Item called *my first CI* depending on the Terraform operation being performed (ie **plan**, **apply**, **destroy**).  "My first CI" has a number of attributes such as *short_description*,
 *ip_address* and so on, all of which are also updated or added depending on the Terraform operation.

The ServiceNow Provider requires three parameters to enable the connection to ServiceNow from the Provider:
1.  The URL to connect to eg https://<domain>.servicenow.com/
2.  The userid to connect with.
3.  The password to use.

The userid and password used must have the required privileges in the ServiceNow CMDB to perform the operations carried out by Terraform.  Typically this is a user with the **cmdb admin** role.

The line in the example below starting with **resource** is the Terraform resource.  Using the example below:
1.  The **resource** is a member of the *servicenowcmdb* provider because the resource specifier starts with *servicenowcmdb*.
2.  The name of the Configuration Item Class is *windows_server*.  This is specified by appending the name of the Configuration Item to the Provider, in this case it is **servicenowcmdb_windows_server**.
3.  The identifier assigned to this resource is **testci4** and this is used to identify this resource in the terraform database / configuration file.
4.  The name of the Windows Server Configuration Item will be "my first CI".  This can be identified by the "name" attribute on the **resource**.
5.  The remaining lines define additional attributes for "my first CI".  These attributes correspond to the attributes for the Configuration Item in the ServiceNow CMDB.
6.  The location attribute is interesting because this refers to a ServiceNow **reference** field which need three nested attributes to be represented correctly, more on this later.

This is a trivial example, better examples using conditional expressions and reference fields from other providers will be defined in the tutorial to be developed soon!!!
```
provider "servicenowcmdb" {
  instance_url = "${var.ServiceNowUrl}"
  username     = "${var.ServiceNowUser}"
  password     = "${var.ServiceNowPwd}"
}

variable "ServiceNowUrl"  {}
variable "ServiceNowUser" {}
variable "ServiceNowPwd"  {}

resource "servicenowcmdb_windows_server" "testci4" {
    name              	= "my first CI"
    short_description   = "and now for something completely stuffed this is from terraform"
    ip_address          = "10.0.0.2"
    discovery_source    = "Terraform"
    operational_status  = "1"
    install_status      = "1"
    location            = {
        value = "25aba2870a0a0bb300926e24d3a5528c"
		link = "https://dev75324.service-now.com//api/now/table/cmn_location/25aba2870a0a0bb300926e24d3a5528c"
		display_value = "1112 18th Street, Plano,TX"
	}
}
```
Using the Provider
--
This Provider is a bit different to existing Terraform Providers because the ServiceNow CMDB is highly configurable (ie customers can add and change CIs in their ServiceNow systems) plus the CMDB is constantly evolving
as new versions of ServiceNow are released.

#### How is this Provider Different?
In order to accommodate the dynamic nature of the ServiceNow CMDB, the provider also needs to be dynamic.  To achieve this, a utility program called "generatesource" is provided which dynamically generates the code required to build the Terraform Provider to talk to ServiceNow.  This utility uses the ServiceNow CMDB API to
connect to the server and retrieve the CMDB metadata for each Configuration Item Class.  It then generates the necessary source code to perform the Create, Read, Update and Delete (CRUD) operations 
for each Terraform resource.

The steps required to use this provider are:

1. Install GO on your workstation or server
2. Setup the "GOPATH" variable (eg /users/philw/go/), if you are having problems building or using the provider, a mis-configured "GOPATH" is normally the cause.
3. clone this repository to the src directory within your "GOPATH" eg /users/philw/go/src/terraform-provider-servicenowcmdb **NOTE** The name of the subdirectory *terraform-provider-servicenowcmdb* is important and cannot be change.
4. change in /users/philw/go/src/terraform-provider-servicenowcmdb/servicenowcmdb/generateprovidersource
5. Setup the following environment variables to enable the utility to log into ServiceNow:
    * ServiceNowUser
    * ServiceNowPwd
    * ServiceNowUrl
6. Run go mod init.  This is required to enable GO to download the required GO modules used by the utility.
7. Run go Mod download
8. Run go build
9. Run generateprovidersource -c cmdb_ci_server:recurse
10. Change back to directory /users/philw/go/src/terraform-provider-servicenowcmdb
11. run go build
12. run go install

And that's it.  The Terraform provider for ServiceNow has now been custom built and installed.  It is configured to 
enable the use of all Configuration Item Classes that are related to the *cmdb_ci_server* class. (eg "Windows Server" etc).

Now copy the example Terraform HCL script defined as the example above and run it using Terraform (assuming you have Terraform installed, if not 
then go to the Terraform web site to read how to download and install, it's pretty simple).  To run this script do the following:

1.  terraform init
```
PS C:\Users\phil\temp\xx> terraform init

Initializing the backend...

Initializing provider plugins...

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```
2.  run the command **terraform plan** and the following output is displayed
```
PS C:\Users\phil\temp\xx> terraform plan
var.ServiceNowPwd
  Enter a value: XXXXXXX

var.ServiceNowUrl
  Enter a value: https://XXXXXX.service-now.com/

var.ServiceNowUser
  Enter a value: admin

Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.

servicenowcmdb_windows_server.testci4: Refreshing state... [id=868b30562f21c410dadacfedf699b6ee]

------------------------------------------------------------------------

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  ~ update in-place

Terraform will perform the following actions:

  # servicenowcmdb_windows_server.testci4 will be updated in-place
  ~ resource "servicenowcmdb_windows_server" "testci4" {
        asset                = {}
        assigned_to          = {}
        assignment_group     = {}
        attributes           = "test"
        can_print            = "false"
        category             = "Hardware"
        cd_rom               = "false"
        change_control       = {}
        company              = {
            "display_value" = "ACME Australia"
            "link"          = "https://dev75324.service-now.com//api/now/table/core_company/df7d53303710200044e0bfc8bcbe5dac"
            "value"         = "df7d53303710200044e0bfc8bcbe5dac"
        }
        cost_center          = {}
        cpu_manufacturer     = {
            "display_value" = "Acer"
            "link"          = "https://dev75324.service-now.com//api/now/table/core_company/820351a1c0a8018b67c73d51c074097c"
            "value"         = "820351a1c0a8018b67c73d51c074097c"
        }
        department           = {}
        discovery_source     = "Terraform"
        dr_backup            = {}
        duplicate_of         = {}
        first_discovered     = "2019-12-14 22:11:40"
        id                   = "868b30562f21c410dadacfedf699b6ee"
        install_status       = "1"
      ~ ip_address           = "10.0.0.3" -> "10.0.0.4"
        last_discovered      = "2019-12-17 05:14:21"
        location             = {
            "display_value" = "1112 18th Street, Plano,TX"
            "link"          = "https://dev75324.service-now.com//api/now/table/cmn_location/25aba2870a0a0bb300926e24d3a5528c"
            "value"         = "25aba2870a0a0bb300926e24d3a5528c"
        }
        maintenance_schedule = {}
        managed_by           = {}
        manufacturer         = {
            "display_value" = "ACME Switzerland"
            "link"          = "https://dev75324.service-now.com//api/now/table/core_company/8ec33b193790200044e0bfc8bcbe5d35"
            "value"         = "8ec33b193790200044e0bfc8bcbe5d35"
        }
        model_id             = {}
        monitor              = "false"
        name                 = "My server 4"
        operational_status   = "1"
        owned_by             = {
            "display_value" = "Abraham Lincoln"
            "link"          = "https://dev75324.service-now.com//api/now/table/sys_user/a8f98bb0eb32010045e1a5115206fe3a"
            "value"         = "a8f98bb0eb32010045e1a5115206fe3a"
        }
        schedule             = {}
        short_description    = "and now for something completely stuffed this is from terraform"
        skip_sync            = "false"
        subcategory          = "Computer"
        support_group        = {}
        supported_by         = {}
        sys_mod_count        = "14"
        sys_updated_by       = "admin"
        sys_updated_on       = "2020-01-04 06:21:52"
        unverified           = "false"
        vendor               = {}
        virtual              = "false"
    }

Plan: 0 to add, 1 to change, 0 to destroy.

------------------------------------------------------------------------

Note: You didn't specify an "-out" parameter to save this plan, so Terraform
can't guarantee that exactly these actions will be performed if
"terraform apply" is subsequently run.
```
3. execute **terraform apply**
```
PS C:\Users\phil\temp\xx> terraform apply
var.ServiceNowPwd
  Enter a value: XXXXXX

var.ServiceNowUrl
  Enter a value: https://XXXXXXX.service-now.com/

var.ServiceNowUser
  Enter a value: admin

servicenowcmdb_windows_server.testci4: Refreshing state... [id=868b30562f21c410dadacfedf699b6ee]

An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  ~ update in-place

Terraform will perform the following actions:

  # servicenowcmdb_windows_server.testci4 will be updated in-place
  ~ resource "servicenowcmdb_windows_server" "testci4" {
        asset                = {}
        assigned_to          = {}
        assignment_group     = {}
        attributes           = "test"
        can_print            = "false"
        category             = "Hardware"
        cd_rom               = "false"
        change_control       = {}
        company              = {
            "display_value" = "ACME Australia"
            "link"          = "https://dev75324.service-now.com//api/now/table/core_company/df7d53303710200044e0bfc8bcbe5dac"
            "value"         = "df7d53303710200044e0bfc8bcbe5dac"
        }
        cost_center          = {}
        cpu_manufacturer     = {
            "display_value" = "Acer"
            "link"          = "https://dev75324.service-now.com//api/now/table/core_company/820351a1c0a8018b67c73d51c074097c"
            "value"         = "820351a1c0a8018b67c73d51c074097c"
        }
        department           = {}
        discovery_source     = "Terraform"
        dr_backup            = {}
        duplicate_of         = {}
        first_discovered     = "2019-12-14 22:11:40"
        id                   = "868b30562f21c410dadacfedf699b6ee"
        install_status       = "1"
      ~ ip_address           = "10.0.0.3" -> "10.0.0.4"
        last_discovered      = "2019-12-17 05:14:21"
        location             = {
            "display_value" = "1112 18th Street, Plano,TX"
            "link"          = "https://dev75324.service-now.com//api/now/table/cmn_location/25aba2870a0a0bb300926e24d3a5528c"
            "value"         = "25aba2870a0a0bb300926e24d3a5528c"
        }
        maintenance_schedule = {}
        managed_by           = {}
        manufacturer         = {
            "display_value" = "ACME Switzerland"
            "link"          = "https://dev75324.service-now.com//api/now/table/core_company/8ec33b193790200044e0bfc8bcbe5d35"
            "value"         = "8ec33b193790200044e0bfc8bcbe5d35"
        }
        model_id             = {}
        monitor              = "false"
        name                 = "My server 4"
        operational_status   = "1"
        owned_by             = {
            "display_value" = "Abraham Lincoln"
            "link"          = "https://dev75324.service-now.com//api/now/table/sys_user/a8f98bb0eb32010045e1a5115206fe3a"
            "value"         = "a8f98bb0eb32010045e1a5115206fe3a"
        }
        schedule             = {}
        short_description    = "and now for something completely stuffed this is from terraform"
        skip_sync            = "false"
        subcategory          = "Computer"
        support_group        = {}
        supported_by         = {}
        sys_mod_count        = "14"
        sys_updated_by       = "admin"
        sys_updated_on       = "2020-01-04 06:21:52"
        unverified           = "false"
        vendor               = {}
        virtual              = "false"
    }

Plan: 0 to add, 1 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

servicenowcmdb_windows_server.testci4: Modifying... [id=868b30562f21c410dadacfedf699b6ee]
servicenowcmdb_windows_server.testci4: Modifications complete after 2s [id=868b30562f21c410dadacfedf699b6ee]

Apply complete! Resources: 0 added, 1 changed, 0 destroyed.
```

And that's it. NOTE: Once you have generated the provider using the **generateprovidersource** utility, you do not have to do ths again
unless you modify the CMDB in ServiceNow to have new Configuration Item classes that you want to include or change attributes on 
Configuration Item Classes, for example adding an attribute called "Service Owner" to the **cmdb_ci_server** class would require 
the utility to be executed again if you wanted to be able to update the "Service Owner" attribute from the Terraform Provider.

Next Steps:
This provider is by no mean Production ready.  It is a minimum viable product to prove the capability.  The following development
is still required:

1.  Add Data Providers : This will enable the Terraform Provider to be able to lookup specific database tables in ServiceNow
to be used in the HCL scripts.  For example having the ability to lookup a location reference in the ServiceNow Location table
and then using that information when building infrastructure (eg using tags) and updating a configuration item in the CMDB.

2. CI Relationships:  Having the ability to create relationships between Configuration Items will enable you to relate Configuration
Items to each other that adds business specific configuration to your CMDB that cannot be discovered using automated processes.

3. Code optimisation: The code is currently a bit smelly as I'm learning GO and Terraform at the same time, it has been an
interesting exercise :-)
