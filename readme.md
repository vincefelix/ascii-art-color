**Hi Talent🤗️,**

🧑‍🏫️*Fonctionnement:*
Le programme ascii-art-justify est une continuité des programmes ascii-art précédent (Fs, Output, ascii-art) auxquelles elle vient ajouter la fonctionnalité de centrer, aligner à droite, à gauche ou justifier une représentation ascii art graphique.
	
	
	Il se référe à un flag entré en argument afin de connaitre l'action à effectuer sur la représentation ascii de celui-ci.
	Le flag devra avoir la syntaxe suivante "--align=" avec un type correspondant à l'action à mener soit center, left, right, justify pour respectivement centrer, aligner à gauche, aligner à droite ou justifier la représentation.


Afin de mener à bien cette tache nous avons récupéré la taille du terminal grâce à "stty size" puis jouer sur un calcul des espaces restants ou à combler en les (espaces) répartissant à gauche, à droite ou de part et d'autre de la représentation graphique de sorte à ce que toute la taille du terminal soit occupée.

	Usage: go run . [OPTION] [STRING] [BANNER]

	Example: go run . --align=right something standard


		 ____                                              _   _   _           _
		|  _ \                                            | | (_) | |         | | 
		| |_) |   ___    _ __           __ _   _   _    __| |  _  | |_        | | 
		|  _ <   / _ \  | '_ \         / _` | | | | |  / _` | | | | __|       | | 
		| |_) | | (_) | | | | |       | (_| | | |_| | | (_| | | | \ |_        |_| 
		|____/   \___/  |_| |_|        \__,_|  \__,_|  \__,_| |_|  \__|       (_) 
                                                                          
                                                                          