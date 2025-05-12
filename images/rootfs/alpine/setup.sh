#!/bin/sh

SERIAL_PORT='ttyS0'

rc_add() {
	local runlevel="$1"; shift  # runlevel name
	local services="$*"  # names of services

	local svc; for svc in $services; do
		mkdir -p etc/runlevels/$runlevel
		ln -s /etc/init.d/$svc etc/runlevels/$runlevel/$svc
		echo " * service $svc added to runlevel $runlevel"
	done
}

#-----------------------------------------------------------------------
einfo 'Configuring system'

# We just prepare this config, but don't start the networking service.
cat > etc/network/interfaces <<-EOF
	auto eth0
	iface eth0 inet dhcp

	post-up /etc/network/if-post-up.d/*
	post-down /etc/network/if-post-down.d/*
EOF

mkdir -p etc/network/if-post-up.d
mkdir -p etc/network/if-post-down.d

if [ "$SERIAL_PORT" ]; then
	echo "$SERIAL_PORT" >> etc/securetty
	sed -Ei "s|^[# ]*($SERIAL_PORT:.*)|\1|" etc/inittab
fi

#-----------------------------------------------------------------------
einfo 'Enabling base system services'

rc_add sysinit devfs dmesg mdev hwdrivers
[ -e etc/init.d/cgroups ] && rc_add sysinit cgroups ||:  # since v3.8

rc_add boot modules hwclock swap hostname sysctl bootmisc syslog
# urandom was renamed to seedrng in v3.17
[ -e etc/init.d/seedrng ] && rc_add boot seedrng || rc_add boot urandom

rc_add shutdown killprocs savecache mount-ro
