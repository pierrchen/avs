package specconv

// DefaultSpec is the default config.json will be created when calling avs init
var DefaultSpec = `
{
    "version": {
        "schema": "0.1",
        "android": "Android O"
    },
    "product": {
        "name": "override by avs s",
        "device": "override by avs s",
        "brand": "override by avs s",
        "model": "override by avs s",
        "manufacture": "override by avs s",
        "inherit_products": [
            "full_base"
        ]
    },
    "boardConfig": {
        "partition_table": {
            "flash_block_size": "512",
            "scheme": "mbr",
            "partitions": [
                {
                    "name": "system",
                    "type": "ext4",
                    "size": "1073741312"
                },
                {
                    "name": "userdata",
                    "type": "ext4",
                    "size": "5456789504"
                },
                {
                    "name": "cache",
                    "type": "ext4",
                    "size": "1073741312"
                }
            ]
        },
        "target": {
            "archs": [
                {
                    "name": "arm",
                    "variant": "armv7-a-neon",
                    "cpu": {
                        "variant": "cortex-a9",
                        "abi": "armeabi-v7a",
                        "abi2": "armeabi"
                    }
                }
            ],
            "no_recovery": true,
            "no_radioimage": true,
            "binder": "64"
        },
        "bootloader": {
            "has_2nd_bootloader": true
        },
        "selinux": {
            "mode": "permissive",
            "policyDir": "auto set when avs s"
        },
        "board_features": [
            "android.software.app_widgets.xml",
            "android.hardware.screen.landscape.xml",
            "android.hardware.usb.host.xml",
            "android.software.print.xml",
            "android.software.webview.xml",
            "android.hardware.ethernet.xml"
        ]
    },
    "boot_image": {
        "kernel": {
            "cmd_line": "TODO: set me here",
            "local_kernel": "auto set when avs s",
            "local_dtb": "auto set when avs s"
        },
        "rootfs_overlay": {
            "fstab": {
                "name": "fstab.default",
                "mounts": [
                    {
                        "src": "TODO:/dev/block/mmcblk0p1 ",
                        "dst": "/system",
                        "type": "ext4",
                        "mnt_flag": "ro,barrier=1,noatime",
                        "fs_mgr_flag": "wait"
                    },
                    {
                        "src": "TODO:/dev/block/mmcblk0p2",
                        "dst": "/cache",
                        "type": "ext4",
                        "mnt_flag": "nosuid,nodev,noatime,barrier=1,data=ordered",
                        "fs_mgr_flag": "wait"
                    },
                    {
                        "src": "TODO:/dev/block/mmcblk0p3",
                        "dst": "/data",
                        "type": "ext4",
                        "mnt_flag": "nosuid,nodev,noatime,barrier=1,data=ordered",
                        "fs_mgr_flag": "wait"
                    }
                ]
            },
            "init.rc": [
                {
                    "name": "init.rc.gen",
                    "imports": [
                        "init.usb.rc"
                    ],
                    "actions": [
                        {
                            "triggers": "boot",
                            "commands": [
                                "mount debugfs /sys/kernel/debug /sys/kernel/debug"
                            ]
                        }
                    ]
                }
            ],
            "uevent.rc": {
                "name": "uevent.rc.gen",
                "rules": [
                    {
                        "node": "/dev/tty",
                        "mode": "0666",
                        "uid": "root",
                        "guid": "root"
                    },
                    {
                        "node": "/sys/devices/virtual/input/input*",
                        "attr": "enable",
                        "mode": "0666",
                        "uid": "root",
                        "guid": "input"
                    }
                ]
            }
        }
    },
    "framework_configs": {
        "properties": [
            "dalvik.vm.heapstartsize=5m",
            "dalvik.vm.heapgrowthlimit=96m",
            "dalvik.vm.heapsize=256m",
            "dalvik.vm.heaptargetutilization=0.75",
            "dalvik.vm.heapminfree=512k",
            "dalvik.vm.heapmaxfree=2m"
        ]
    },
    "hals": [
        {
            "name": "keymaster",
            "required_packages": {
                "build": [
                    "android.hardware.keymaster@3.0-impl",
                    "android.hardware.keymaster@3.0-service"
                ]
            }
        },
        {
            "name": "graphics",
            "manifests": [
                {
                    "name": "android.hardware.graphics.composer",
                    "format": "hidl",
                    "transport": {
                        "mode": "hwbinder"
                    },
                    "version": "2.1",
                    "interface": {
                        "name": "IComposer",
                        "instance": "default"
                    }
                },
                {
                    "name": "android.hardware.graphics.allocator",
                    "format": "hidl",
                    "transport": {
                        "mode": "hwbinder"
                    },
                    "version": "2.0",
                    "interface": {
                        "name": "IAllocator",
                        "instance": "default"
                    }
                },
                {
                    "name": "android.hardware.graphics.mapper",
                    "format": "hidl",
                    "transport": {
                        "arch": "32+64",
                        "mode": "passthrough"
                    },
                    "impl": {
                        "level": "generic"
                    },
                    "version": "2.0",
                    "interface": {
                        "name": "IMapper",
                        "instance": "default"
                    }
                }
            ],
            "build_configs": [
                "TARGET_USES_HWC2 := true"
            ],
            "required_packages": {
                "build": [
                    "libion",
                    "android.hardware.graphics.mapper@2.0",
                    "android.hardware.graphics.mapper@2.0-impl",
                    "android.hardware.graphics.allocator@2.0",
                    "android.hardware.graphics.allocator@2.0-impl",
                    "android.hardware.graphics.allocator@2.0-service",
                    "android.hardware.graphics.composer@2.1",
                    "android.hardware.graphics.composer@2.1-impl",
                    "android.hardware.graphics.composer@2.1-service"
                ]
            }
        },
        {
            "name": "audio",
            "features": [
                "android.hardware.audio.output.xml",
                "android.hardware.audio.low_latency.xml"
            ],
            "required_packages": {
                "build": [
                    "android.hardware.audio@2.0-impl",
                    "android.hardware.audio@2.0-service",
                    "android.hardware.audio.effect@2.0-impl",
                    "audio.a2dp.default",
                    "audio.usb.default",
                    "audio.r_submix.default",
                    "audio.primary.poplar:v"
                ]
            },
            "runtime_configs": [
                {
                    "src": "TODO:$(LOCAL_PATH)/audio_policy.conf"
                }
            ]
        },
        {
            "name": "media.codec",
            "runtime_configs": [
                {
                    "src": "TODO:$(LOCAL_PATH)/media_codecs.xml",
                    "destDir": "system/etc/media_codecs.xml"
                },
                {
                    "src": "frameworks/av/media/libstagefright/data/media_codecs_google_video.xml"
                },
                {
                    "src": "frameworks/av/media/libstagefright/data/media_codecs_google_audio.xml"
                }
            ]
        },
        {
            "name": "bluetooth",
            "build_configs": [
                "BOARD_HAVE_BLUETOOTH_BCM := true"
            ]
        }
    ],
    "vendor_raw": {}
}
`
