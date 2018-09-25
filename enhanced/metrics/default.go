// Code generated by go generate; DO NOT EDIT.
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var Default = map[string]Metric{
	"rdsosmetrics_General_engine": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_engine",
			"The database engine for the DB instance.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_General_instanceID": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_instanceID",
			"The DB instance identifier.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_General_instanceResourceID": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_instanceResourceID",
			"A region-unique, immutable identifier for the DB instance, also used as the log stream identifier.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_General_numVCPUs": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_numVCPUs",
			"The number of virtual CPUs for the DB instance.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_General_timestamp": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_timestamp",
			"The time at which the metrics were taken.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_General_uptime": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_uptime",
			"The amount of time that the DB instance has been active.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_General_version": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_General_version",
			"The version of the OS metrics' stream JSON format.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_cpu_average": {
		Desc: prometheus.NewDesc(
			"node_cpu_average",
			"The percentage of CPU utilization. Units: Percent",
			[]string{"instance", "region", "mode"},
			map[string]string{"cpu": "All"},
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_avgQueueLen": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_avgQueueLen",
			"The number of requests waiting in the I/O device's queue. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_avgReqSz": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_avgReqSz",
			"The average request size, in kilobytes. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_await": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_await",
			"The number of milliseconds required to respond to requests, including queue time and service time. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_diskQueueDepth": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_diskQueueDepth",
			"The number of outstanding read/write requests waiting to access the disk.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_readIOsPS": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readIOsPS",
			"The number of read operations per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_disk_bytes_read": {
		Desc: prometheus.NewDesc(
			"node_disk_bytes_read",
			"The total number of bytes read. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"rdsosmetrics_diskIO_readKbPS": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readKbPS",
			"The number of kilobytes read per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_readLatency": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readLatency",
			"The average amount of time taken per disk I/O operation.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_readThroughput": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_readThroughput",
			"The average number of bytes read from disk per second.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_rrqmPS": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_rrqmPS",
			"The number of merged read requests queued per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_tps": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_tps",
			"The number of I/O transactions per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_util": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_util",
			"The percentage of CPU time during which requests were issued. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_writeIOsPS": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeIOsPS",
			"The number of write operations per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_disk_bytes_written": {
		Desc: prometheus.NewDesc(
			"node_disk_bytes_written",
			"The total number of bytes written. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"rdsosmetrics_diskIO_writeKbPS": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeKbPS",
			"The number of kilobytes written per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_writeLatency": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeLatency",
			"The average amount of time taken per disk I/O operation.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_writeThroughput": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_writeThroughput",
			"The average number of bytes written to disk per second.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_diskIO_wrqmPS": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_diskIO_wrqmPS",
			"The number of merged write requests queued per second. This metric is not available for Amazon Aurora.",
			[]string{"instance", "region", "id", "device"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_filesystem_avail": {
		Desc: prometheus.NewDesc(
			"node_filesystem_avail",
			"The amount of disk space available in the file system, in bytes.",
			[]string{"instance", "region", "id", "name", "mountPoint"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"rdsosmetrics_fileSys_maxFiles": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_maxFiles",
			"The maximum number of files that can be created for the file system.",
			[]string{"instance", "region", "id", "name", "mountPoint"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_filesystem_size": {
		Desc: prometheus.NewDesc(
			"node_filesystem_size",
			"The total number of disk space available for the file system, in bytes.",
			[]string{"instance", "region", "id", "name", "mountPoint"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"rdsosmetrics_fileSys_used": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_used",
			"The amount of disk space used by files in the file system, in kilobytes.",
			[]string{"instance", "region", "id", "name", "mountPoint"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_fileSys_usedFilePercent": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_usedFilePercent",
			"The percentage of available files in use.",
			[]string{"instance", "region", "id", "name", "mountPoint"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_fileSys_usedFiles": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_usedFiles",
			"The number of files in the file system.",
			[]string{"instance", "region", "id", "name", "mountPoint"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_fileSys_usedPercent": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_fileSys_usedPercent",
			"The percentage of the file-system disk space in use.",
			[]string{"instance", "region", "id", "name", "mountPoint"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_loadAverageMinute_fifteen": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_loadAverageMinute_fifteen",
			"The number of processes requesting CPU time over the last 15 minutes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_loadAverageMinute_five": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_loadAverageMinute_five",
			"The number of processes requesting CPU time over the last 5 minutes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_load1": {
		Desc: prometheus.NewDesc(
			"node_load1",
			"The number of processes requesting CPU time over the last minute.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_memory_Active": {
		Desc: prometheus.NewDesc(
			"node_memory_Active",
			"The amount of assigned memory, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_memory_Buffers": {
		Desc: prometheus.NewDesc(
			"node_memory_Buffers",
			"The amount of memory used for buffering I/O requests prior to writing to the storage device, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_memory_Cached": {
		Desc: prometheus.NewDesc(
			"node_memory_Cached",
			"The amount of memory used for caching file system–based I/O, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_vmstat_nr_dirty": {
		Desc: prometheus.NewDesc(
			"node_vmstat_nr_dirty",
			"The amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_memory_MemFree": {
		Desc: prometheus.NewDesc(
			"node_memory_MemFree",
			"The amount of unassigned memory, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"rdsosmetrics_memory_hugePagesFree": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesFree",
			"The number of free huge pages.Huge pages are a feature of the Linux kernel.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_memory_hugePagesRsvd": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesRsvd",
			"The number of committed huge pages.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_memory_hugePagesSize": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesSize",
			"The size for each huge pages unit, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_memory_hugePagesSurp": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesSurp",
			"The number of available surplus huge pages over the total.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_memory_hugePagesTotal": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_hugePagesTotal",
			"The total number of huge pages for the system.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_memory_Inactive": {
		Desc: prometheus.NewDesc(
			"node_memory_Inactive",
			"The amount of least-frequently used memory pages, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_memory_Mapped": {
		Desc: prometheus.NewDesc(
			"node_memory_Mapped",
			"The total amount of file-system contents that is memory mapped inside a process address space, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_memory_PageTables": {
		Desc: prometheus.NewDesc(
			"node_memory_PageTables",
			"The amount of memory used by page tables, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_memory_Slab": {
		Desc: prometheus.NewDesc(
			"node_memory_Slab",
			"The amount of reusable kernel data structures, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_memory_MemTotal": {
		Desc: prometheus.NewDesc(
			"node_memory_MemTotal",
			"The total amount of memory, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"rdsosmetrics_memory_writeback": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_memory_writeback",
			"The amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_network_rx": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_network_rx",
			"The number of bytes received per second.",
			[]string{"instance", "region", "id", "interface"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_network_tx": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_network_tx",
			"The number of bytes uploaded per second.",
			[]string{"instance", "region", "id", "interface"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_processList_cpuUsedPc": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_cpuUsedPc",
			"The percentage of CPU used by the process.",
			[]string{"instance", "region", "id", "name"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_processList_id": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_id",
			"The identifier of the process.",
			[]string{"instance", "region", "id", "name"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_processList_memoryUsedPc": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_memoryUsedPc",
			"The amount of memory used by the process, in kilobytes.",
			[]string{"instance", "region", "id", "name"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_processList_parentID": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_parentID",
			"The process identifier for the parent process of the process.",
			[]string{"instance", "region", "id", "name"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_processList_rss": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_rss",
			"The amount of RAM allocated to the process, in kilobytes.",
			[]string{"instance", "region", "id", "name"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_processList_tgid": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_tgid",
			"The thread group identifier, which is a number representing the process ID to which a thread belongs.This identifier is used to group threads from the same process.",
			[]string{"instance", "region", "id", "name"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_processList_vmlimit": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_vmlimit",
			"TODO",
			[]string{"instance", "region", "id", "name"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_processList_vss": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_processList_vss",
			"The amount of virtual memory allocated to the process, in kilobytes.",
			[]string{"instance", "region", "id", "name"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_swap_cached": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_swap_cached",
			"The amount of swap memory, in kilobytes, used as cache memory.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_memory_SwapFree": {
		Desc: prometheus.NewDesc(
			"node_memory_SwapFree",
			"The total amount of swap memory free, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_vmstat_pswpin": {
		Desc: prometheus.NewDesc(
			"node_vmstat_pswpin",
			"Number of kilobytes the system has swapped in from disk per second (disk reads).",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_vmstat_pswpout": {
		Desc: prometheus.NewDesc(
			"node_vmstat_pswpout",
			"Number of kilobytes the system has swapped out to disk per second (disk writes).",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_memory_SwapTotal": {
		Desc: prometheus.NewDesc(
			"node_memory_SwapTotal",
			"The total amount of swap memory available, in bytes.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1024,
	},
	"node_procs_blocked": {
		Desc: prometheus.NewDesc(
			"node_procs_blocked",
			"The number of tasks that are blocked.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"node_procs_running": {
		Desc: prometheus.NewDesc(
			"node_procs_running",
			"The number of tasks that are running.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_tasks_sleeping": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_sleeping",
			"The number of tasks that are sleeping.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_tasks_stopped": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_stopped",
			"The number of tasks that are stopped.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_tasks_total": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_total",
			"The total number of tasks.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
	"rdsosmetrics_tasks_zombie": {
		Desc: prometheus.NewDesc(
			"rdsosmetrics_tasks_zombie",
			"The number of child tasks that are inactive with an active parent task.",
			[]string{"instance", "region"},
			map[string]string(nil),
		),
		Unit: 1,
	},
}