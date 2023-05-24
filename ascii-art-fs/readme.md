**Hi Talent🤗️,**

🧑‍🏫️*Fonctionnement:*
Le programme ascii-art-fs recoit une string puis travaille avec la fonction main👨‍💻️ en utilisant os.Args pour la récupération de ces arguments et en ajoutant au banner une extension ".txt" afin d'utiliser les représentations graphiques adéquates.
Utilisation d'un tableau à double dimension permettant de situer précisemment et de stocker la représentation graphique correspondante après avoir splité par '\n'.
- Si un caractere n'a pas sa représentation graphique la fonction affiche un message d'erreur de "Non affichable, aucune correspondance graphique"
- Sinon la fonction affiche la représentation graphique adaptée en usant du banner de la string correspondante.
	
*Usage :*
- le programme s'exécute avec la commange go run . [string] [banner] et pour afficher avec les détails on peut piper par | cat -e "(💬️string à afficher)" Exemple: go run . "Manger" shadow | cat -e -> représentation graphique de Manger avec shadow
- Dans le cas ou la string à afficher ne contient qu'une ponctuation au lieu des doubles cotes, utiliser les apostrophes pour rentrer la string à run (💬️Ex: go run . '.' thinkertoy )
- NB: Les retours à la ligne dans ce programme sont matérialisés par un '\n'
							
						 ____                                              _   _   _           _
						|  _ \                                            | | (_) | |         | | 
						| |_) |   ___    _ __           __ _   _   _    __| |  _  | |_        | | 
						|  _ <   / _ \  | '_ \         / _` | | | | |  / _` | | | | __|       | | 
						| |_) | | (_) | | | | |       | (_| | | |_| | | (_| | | | \ |_        |_| 
						|____/   \___/  |_| |_|        \__,_|  \__,_|  \__,_| |_|  \__|       (_) 
                                                                          
                                                                          
