GOOF----LE-4-2.0¬&      ] g 4       h      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  language¤	g  assembly¤	 ¤		g  filenameS¤	
f  language/assembly.scm¤	g  importsS¤	g  rnrs¤	g  bytevectors¤	 ¤	 ¤	g  system¤	g  base¤	g  pmatch¤	 ¤	 ¤	g  vm¤	g  instruction¤	 ¤	 ¤	g  srfi¤	g  srfi-1¤	 ¤	g  selectS¤	g  fold¤	 ¤	 ¤	  ¤	!g  exportsS¤	"g  byte-length¤	#g  addr+¤	$g  align-program¤	%g  
align-code¤	&g  align-block¤	'g  assembly-pack¤	(g  assembly-unpack¤	)g  object->assembly¤	*g  assembly->object¤	+"#$%&'()* 	¤	,g  set-current-module¤	-, ¤	., ¤	/g  *program-header-len*¤	0g  	*len-len*¤	1g  load-program¤	2g  error¤	3f  unknown instruction¤	4g  
load-array¤	5g  bytevector-length¤	6g  load-symbol¤	7g  string-length¤	8g  load-wide-string¤	9g  load-string¤	:g  load-number¤	;g  instruction-length¤	<g  *program-alignment*¤	=g  	make-list¤	>g  nop¤	?> ¤	@g  code-alignment¤	Ag  append¤	BA ¤	CA ¤	Dg  	make-int8¤	ED
 ¤	Fg  make-int8:0¤	GEF ¤	HD ¤	Ig  make-int8:1¤	JHI ¤	KGJ ¤	Lg  *abbreviations*¤	Mg  map¤	Ng  *expansions*¤	Og  	assoc-ref¤	Pg  	make-true¤	QP ¤	Rg  
make-false¤	SR ¤	Tg  make-nil¤	UT ¤	Vg  make-eol¤	WV ¤	Xg  char?¤	Yg  char->integer¤	Zg  
make-char8¤	[g  make-char32¤	\g  integer?¤	]g  exact?¤	^g  
make-int64¤	_g  bytevector->u8-list¤	`g  make-bytevector¤	ag  big¤	bg  make-uint64¤	cg  
make-int16¤	dg  string->symbol¤	eg  integer->char¤	fg  u8-list->bytevector¤C 5       h  ê   ]4	
 !+5 4. >  "  G   	/R	0R1/"234056789:;  hP  Ä  ]1" " Ê" z" '"  ×"   $  ~  &  d$  U$  =$  $$  45"  
C 6 6 6 6
C $  A  &  +$   (  45C"ÿÿ9"ÿÿ5"ÿÿ1"ÿÿ- $  A  	&  +$   (  4
5C"ÿþé"ÿþå"ÿþá"ÿþÝ $  D  &  .$  #(  	4
5C"ÿþ"ÿþ"ÿþ"ÿþ $  A  &  +$   (  4
5C"ÿþF"ÿþB"ÿþ>"ÿþ: $  A  &  +$   (  4
5C"ÿýö"ÿýò"ÿýî"ÿýê $  #  45
$  	45C"ÿý¼"ÿý¸  ¼      g  assembly
	N g  vx	'  g  vy		'  g  vy		=  g  vx		P  g  vy		P  g  vx		^	 g  vx ® é g  vy	 ® é g  vx	 Ç á g  vy	 Ç á g  vx ú5 g  vy	 ú5 g  vx	- g  vy	- g  vxF g  vy	F g  vx	_| g  vy	_| g  vxÐ g  vy	Ð g  vx	®È g  vy	®È g  vxá g  vy	á g  vx	ú g  vy	ú g  vx*J  g  filenamef  language/assembly.scm
	%
		&		g	4		p	4	#	q	4	0	x	4	,	~	4	 	7	 	7	
 	7	 	7	
 	7	 	7	
 	7	 	7	
 ¡	&	 Ô	2	 Õ	2	 Ü	2	 á	&	 	0	!	0	(	0	-	&	l	.	o	.	v	.	w	.	|	&	»	,	¼	,	Ã	,	È	&		*		*		*		&	1	'	9	'	=	&	>	(	
E	(	J	&	 ,	N  g  nameg  byte-length C"R	<R"   h   y   ]4 5C   q       g  x
		 g  len		  g  filenamef  language/assembly.scm
	=			=			=	 			   C    h      ] 6    ~       g  addr
		 g  code		  g  filenamef  language/assembly.scm
	<
		=	 			  g  nameg  addr+ C#R=?       h   Ô   ] 6  Ì       g  addr
		 g  	alignment		 g  
header-len			  g  filenamef  language/assembly.scm
	A
		C	 		C			B			B			E			B	 			  g  nameg  code-alignment C@R       h   z   ]C    r       g  addr
		  g  filenamef  language/assembly.scm
	G
		H	 		  g  nameg  align-block C&RC@   h   Â   ]45  6  º       g  code
		 g  addr		 g  	alignment			 g  
header-len			  g  filenamef  language/assembly.scm
	J
		K			K	 			  g  nameg  
align-code C%R%<   h      ] 6          g  prog
		 g  addr		  g  filenamef  language/assembly.scm
	N
		O	 			  g  nameg  align-program C$RKLR4Mi   h   p   ]  C      h       g  x
		
  g  filenamef  language/assembly.scm
	Z			Z			Z	!			Z	 		
   CLi5NROL       h       ]	4 5$  C C              g  code
		 g  t		  g  filenamef  language/assembly.scm
	\
		]			]	 		  g  nameg  assembly-pack C'RON   h       ]	4 5$  C C              g  code
		 g  t		  g  filenamef  language/assembly.scm
	`
		a			a	 		  g  nameg  assembly-unpack C(RQSUWXYZ[\]^_`abc'D h    ]	 &  C &  C &  C (  C"  64 5$  )4 5 ÿ$  4 5 C4 5 CC4	 5$ #4
 5$ "  O        
$  ÿÿÿÿÿÿÿ"  $  44	5
 Å5CC"  ?
 $  3 ÿÿÿÿÿÿÿÿ$  44	5
 Ä5C"ÿÿr"ÿÿn"  P  $  B        $  . 
$          "      C"ÿÿq"ÿÿm	 $    $     6"ÿÿ"ÿÿ"ÿþ"ÿþ         g  x
	 g  b  § g  bv · Æ g  bv ô g  nG[  g  filenamef  language/assembly.scm
	i
			j			j			j			k			j			l		&	j		(	m		.				8	j		9 		C 		G 			I 		J 		S 		V 		W 		` 		c	j		d	n		n	j		o	n		y	n		 	y	 	y	 «	o		 ­	z	 ®	z	 ±	{	* ·	{	  À	|	> Á	|	" È	z	 É	z	 Ì	o		 Ô	t	 Ø	o		 ä	t	 ê	u	 ë	u	  î	v	+ ô	v	! ý	w	? þ	w	#	u	 	u		o			q		o		)	q	#-	q	1	r	5	r	A	r	%G	r	K	s	Q	s	 W	s	2Z	s	c	o		h	o	l	o		r	o	!v	o	z	p		p	,	p		p	 G	  g  nameg  object->assembly C)R6d9[eZ^fabcDVTRP   h  <  ]" Ð" " n" =" å" a" m" y" /"  "  > $  5  &  !$  (  6CCCC $  8  &  "$  (  C"ÿÿ"ÿÿ"ÿÿ"ÿÿ $    &  $  z$  e$  P$  ;
	
(  '               	6"ÿþí"ÿþé"ÿþå"ÿþá"ÿþÝ"ÿþÙ"ÿþÕ $  ;  &  %$  (  6"ÿþ"ÿþ"ÿþ"ÿþ $  å  &  Ï$  Ä$  ¯$  $  
	
$  p

$  [$  F$  1(  4	 5
	³C"ÿý¿"ÿý»"ÿý·"ÿý³"ÿý¯"ÿý«"ÿý§"ÿý£"ÿý"ÿý"ÿý $  å  
&  Ï$  Ä$  ¯$  $  
	
$  p

$  [$  F$  1(  4	 5
	²C"ÿüË"ÿüÇ"ÿüÃ"ÿü¿"ÿü»"ÿü·"ÿü³"ÿü¯"ÿü«"ÿü§"ÿü£ $  u  &  _$  T$  ?(  +        $  C       C"ÿü/"ÿü+"ÿü'"ÿü#"ÿü $  I  &  3$  ((   $  C C"ÿûÓ"ÿûÏ"ÿûË"ÿûÇ $  "  &  (  C"ÿû"ÿû"ÿû $  "  &  (  C"ÿûm"ÿûi"ÿûe $  "  &  (  C"ÿû<"ÿû8"ÿû4 $  "  &  (  C"ÿû"ÿû"ÿû     4      g  code
	 g  vx	;	j g  vy		;	j g  vx		T	f g  vy		T	f g  vx	y « g  vy		y « g  vx	  £ g  vy	  £ g  vx ¼Q g  vy	 ¼Q g  vx	 ÕI g  vy	 ÕI g  vx	 æE g  vy	 æE g  vx	 ÷A g  vy	 ÷A g  vx		= g  vy	
= g  vxb g  vy	b g  vx	{ g  vy	{ g  vx¨ g  vy	¨ g  vx	Á g  vy	Á g  vx	Ò{ g  vy	Ò{ g  vx	ãw g  vy	ãw g  vx		ôs g  vy	
ôs g  vx	o g  vy	o g  vx	k g  vy	k g  vx	'g g  vy	'g g  vx	8c g  vy	8c g  vxw g  vy	w g  vx	±o g  vy	±o g  vx	Âk g  vy	Âk g  vx	Óg g  vy	Óg g  vx		äc g  vy	
äc g  vx	õ_ g  vy	õ_ g  vx	[ g  vy	[ g  vx	W g  vy	W g  vx	(S g  vy	(S g  vx÷ g  vy	÷ g  vx	¡ï g  vy	¡ï g  vx	²ë g  vy	²ë g  n	Åç g  vxK g  vy	K g  vx	!C g  vy	!C g  vx\x g  vy	\x g  vx¥ g  vy	¥ g  vx¶Ò g  vy	¶Ò g  vxãÿ g  vy	ãÿ  Ng  filenamef  language/assembly.scm
 
	 		d £		l 	  	, 	- 	3  	4 	9 	= 	 	 	C 	X 	Z 	] 	^ 	c 	3 	H 	J 	M 	N 	S 	Â 	Å 	Å 	Ó 	× 	æ 	ë 	1 		5 	> 	C 	o 	t 	 '	  g  nameg  assembly->object C*RC       â       g  m
		,  g  filenamef  language/assembly.scm		
	.			1	
	6	"
w	%
|	:
Á	<
Ç	A
	]	G

M	J
	N
	V		U
	Z	¡	Y
m	\
7	`
	i
 
 	
   C6 