# xultybau

> > xulta
>
> x1 is beyond/beside/abstract/meta (older sense) to x2 so as to be an abstraction thereof/external thereto, being so beyond/abstract in property/aspect x3, and having features x4, governed by/considered in system/considerations/epistemology/theory x5; x1 is conceptually beyond/external to/higher up than x2 (in, for example, existential or consideration precedence)

> > bangu
>
> x1 is a/the language/dialect used by x2 to express/communicate x3 (si'o/du'u, not quote).

This is going to be a large project. What I want to do here is to be able to 
break sentences into the entire set of their grammatical features for the
purposes of analysis, compilation and execution. In essence, I want a heckin
API that will take a Toki Pona sentence like `pilin mi la ken la ona li toki e ni la jan lawa li pakala e ona` and get out
something like:

```lisp
(sentence
  (context
    (context (pilin (mi)) (ken))
    (subject (ona))
    (verb (toki))
    (object (ni))
  (subject (jan (lawa)))
  (verb (pakala))
  (object (ona)))
```

## Aesthetic

Last night I was walking through the harbor,  
Where the fishing boats are lying on the shore,  
The news had traveled fast and everyone went to be,  
Where the mayor was making a speech,  
And the crowd started cheering,  
When he talked about the glory of it all,  
And the boys  
