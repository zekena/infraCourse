# The backup critera for the working services

## Coverage

 - Webserver with logs
 - Database server through mysql dump
 - DNS master and slave configs
 - Ansible repo will be updated in github asap so no need for other backup
 - Backup server configs will be updated to github.

## Frequency
 All the data that's covered will be backuped every 15 day starting from the 1st day of the month at 1 am.
 Incremential backup will be made every everyday at 12 pm :
 - logs
 - database wordpress
## Usability
  We will use duplicity and cron tabs as our backup system with the help of protocols rsync,file and cp we will be able to complete every task.
## Retention

 - every 1 month webservers logs and dump files will be removed
 - DNS,webserver and mysql config files will be removed every 6 months
 
## Storage

 - backup server
 - backup agent
 - some will be on github
 
