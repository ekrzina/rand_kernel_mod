# Makefile for rngen kernel module

ARCH=arm64
CROSS_COMPILE=aarch64-linux-gnu-
KDIR = $(HOME)/linux

MODULE = rngen

# setting source
obj-m += $(MODULE).o

all:
	$(MAKE) -C $(KDIR) ARCH=$(ARCH) CROSS_COMPILE=$(CROSS_COMPILE) M=$(PWD) modules

clean:
	$(MAKE) -C $(KDIR) ARCH=$(ARCH) CROSS_COMPILE=$(CROSS_COMPILE) M=$(PWD) clean
