# Temporal

## Setup
- https://github.com/temporalio/docker-compose

## Misc
```
Error response from daemon: error while creating mount source path '/host_mnt/Users/thanhnguyen/Desktop/PERSONAL/queue-fluency/golang-temporal/dynamicconfig': mkdir /host_mnt/Users/thanhnguyen/Desktop: operation not permitted
make: *** [up] Error 1
```

- Fix: Grant Docker Full Disk Access:
    - Open System Preferences.
    - Go to Security & Privacy.
    - Select the Privacy tab.
    - Scroll down and select Full Disk Access.
    - Click the lock icon to make changes and add Docker Desktop to the list.