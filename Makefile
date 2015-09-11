U ?= $(shell pwd)

include $(U)/bootstrap.mk
include $(MAKEDIR)/sugarTop.mk

noinst_GOPROGRAMS = \
	linuxProxy

noinst_GOLIBRARIES = \
	nsg \
	edbtypes

nsg_GOSRC = netStackGoUserDefined.go nsgPkg.go
nsg_TYPES = NetStackGo.t

edbtypes_TYPES = EdbTypes.t
edbtypes_GOSRC = EdbTypesRpc.go

linuxProxy_GOSRC = linuxProxy.go
linuxProxy_GOPKGDIRS = .

include $(MAKEDIR)/sugarBottom.mk
