----------
avorf1com 
----------

 
### Installation 
Copier le ficheir avorf1com_[version].tar.gz sur la machine cible dans le repertoire /opt/avorf1com
Exécuter tar xvzf avorf1com_[version].tar.gz

Ce placer dans le repertoire : cd /opt/avorf1com/

> **NOTE:**
>
> - En cas d'une nouvelle installation copier le ficheir avorf1com.supervisor.conf dans le répertoire /etc/supervisor/conf.d
>
> - Copier le fichier /opt/avorf1com/avorf1com_[version]/config.json dans /opt/avorf1com/avorf1com_[version]/config.json
> - Addapter ce fichier en cas de besoin à votre environement cible
> - Copier le fichier /opt/avorf1com/avorf1com_[version]/logger.xml dans /opt/avorf1com/avorf1com_[version]/logger.xml
> - Addapter ce fichier en cas de besoin à votre environement cible


Vérifier si l'application est en cours d'execution via console de supervisor : supervisorctl
Si l'application est en court d'exécution arrêter l'application : stop avorf1com
Quitter le console : exit

Crée un lien symbolic ln -s /opt/avorf1com/avorf1com_[version] /opt/avorf1com/current 
> **NOTE:**
>
> - En cas si le répertoire existe /opt/avorf1com/current. Supprimer rm -rf /opt/avorf1com/current

### Configuraiton
Copier le fichier config.json de répértoire d'instalaltion dans le répértoire de l'application : cp ./samples/config.json config.json
Copier le fichier logger.xml de répértoire d'instalaltion dans le répértoire de l'application : cp ./samples/logger.xml logger.xml


### Mise à jour 
Mise à jour est identique à l'installation sans la partie de la configuration.

Il faut copier le ficheir de la configuration actuel (/opt/avorf1com/current/config.json) dans le repertoire /opt/avorf1com/avorf1com_[version] 

## Notes
		Cluster
		code NODSTO : L'information complimentaire Cluster resource stopped on astnode1
		code NODSTA : L'information complimentaire resource started on astnode1

		Statistiques:
		APPSTA  : application vorimport démarrée
		APPSTO  : application vorimport arrêtée
		MYSQKO  : connexion à la base de données mysql ok
		MYSQOK  : connexion à la base de données mysql ko
		MONGKO  : connexion à la base de données mongo ok
		MONGOK  : connexion à la base de données mongo ko
		TCALOK  : connexion au serveur astersik(l'application) est ok
		TCALKO  : connexion au serveur astersik(l'application) est ko
		CCALOK  : generation d'un appel de test ok 
		CCALKO  : generation d'un appel de test ok 





