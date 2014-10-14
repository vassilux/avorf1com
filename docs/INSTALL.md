----------
avorf1com 
----------

 
### Installation 
Copier le ficheir avorf1com_[version].tar.gz sur la machine cible dans le repertoire /opt/avorf1com
Exécuter tar xvzf avorf1com_[version].tar.gz

Ce placer dans le repertoire : cd /opt/avorf1com/avorf1com_[version]

> **NOTE:**
>
> - En cas d'une nouvelle installation copier le ficheir avorf1com.supervisor.conf dans le répertoire /etc/supervisor/conf.d
>
> - Copier le fichier /opt/avorf1com/avorf1com_[version]/config.sample.json dans /opt/avorf1com/avorf1com_[version]/config.json
> - Addapter ce fichier en cas de besoin à votre environement cible


Vérifier si l'application est en cours d'execution via console de supervisor : supervisorctl
Si l'application est en court d'exécution arrêter l'application : stop avorf1com
Quitter le console : exit

Crée un lien symbolic ln -s /opt/revor/avorf1com_[version] /opt/avorf1com/current 
> **NOTE:**
>
> - En cas si le répertoire existe /opt/avorf1com/current. Supprimer rm -rf /opt/avorf1com/current

### Configuraiton
Copier le fichier config.json de répértoire d'instalaltion dans le répértoire de l'application : cp ./samples/config.json config.json
Copier le fichier logger.xml de répértoire d'instalaltion dans le répértoire de l'application : cp ./samples/logger.xml logger.xml


### Mise à jour 
Mise à jour est identique à l'installation sans la partie de la configuration.

Il faut copier le ficheir de la configuration actuel (/opt/avorf1com/current/config.json) dans le repertoire /opt/avorf1com/avorf1com_[version] 





