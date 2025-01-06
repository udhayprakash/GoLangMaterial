# glog Flags

-----------------------------------------------------------------------------------------------------------------------------
Flag	            Default Value	Description
-----------------------------------------------------------------------------------------------------------------------------
-logtostderr	    false	        Logs are written to standard error instead of files.
                                    go run main.go -logtostderr=true

-alsologtostderr	false	        Logs are written to both files and standard error.
                                    go run main.go -alsologtostderr=true

-stderrthreshold	ERROR	        Logs at or above this threshold are written to stderr (e.g., INFO, WARNING, ERROR, FATAL).
                                    go run main.go -stderrthreshold=WARNING

-log_dir	        "" (empty)	    Directory to write log files. If empty, logs are written to /tmp.
                                    go run main.go -log_dir=./logs

-log_backtrace_at	"" (empty)	    Emit a stack trace when logging at the specified file:line (e.g., main.go:10).
                                    go run main.go -log_backtrace_at=main.go:10

-v	                0	            Set the verbosity level for V-logging. Higher values enable more logs.
                                    go run main.go -v=2

-vmodule	        "" (empty)	    Comma-separated list of pattern=N settings for file-filtered V-logging.
                                    go run main.go -vmodule=main.go=2,utils.go=1

-logtostderr	    false	        Logs are written to standard error instead of files.
-log_file_max_size	1800 (MB)	    Maximum size of a log file before it is rotated (in MB).
                                    go run main.go -log_file_max_size=100

-log_file_max_age	0 (days)	    Maximum age of a log file before it is rotated (in days).
                                    go run main.go -log_file_max_age=7

-log_file_max_count	0	            Maximum number of old log files to retain.
                                    go run main.go -log_file_max_count=5

-log_link	        "" (empty)	    Create a symbolic link to the latest log file with this name.
                                    go run main.go -log_link=latest.log

-----------------------------------------------------------------------------------------------------------------------------


USAGE:

    go run main.go -logtostderr=true -v=2 -log_dir=./logs -stderrthreshold=WARNING
