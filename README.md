# IT Infrastructure services

This is a project for it infrastructure services course, Taltech university,
Cyber-security engineering program
## Getting Started

The EXAM folder has asnible code to deploy services and set up a wordpress instances with MySQL database and Logging & monitoring services with dashborad manager grafana in docker containers.
### Prerequisites

You will need two ubuntu servers connected in one network and you can change the ip address from the host_vars group_vars folders, after setting up the required ip addresses and root access you can move to installing part

### Installing

One command will be enough to make everything work

```
ansible-playbook infra.yaml
```


## Built With

* [Ansible](https://www.ansible.com) - Automation Tool

## Authors


Zeyad Saber Refaei Kenawi [Github account](https://github.com/zekena) 
## Acknowledgments

* Roman kuchin 
* Juri Hudolejev
