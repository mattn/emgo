#define ithead builtin$ItHead

struct minfo {
	builtin$Method;
};

struct tinfo {
	builtin$Type;
	unsafe$Pointer imethods[];
};

typedef struct {
	ithead h$;
	string (*Error)(ival *);
} error;

enum {
	Invalid = 0,
	Bool,
	Int,
	Int8,
	Int16,
	Int32,
	Int64,
	Uint,
	Uint8,
	Uint16,
	Uint32,
	Uint64,
	Uintptr,
	Float32,
	Float64,
	Complex64,
	Complex128,
	Chan,
	Func,
	Interface,
	Map,
	Ptr,
	Slice,
	String,
	Struct,
	UnsafePointer,
	
	Array = -1
};

#define TINFO(i) (((const ithead*)(i).itab$)->typ)

#define IASSIGN(expr, etyp, ityp) INTERFACE(        \
	expr,                                           \
	builtin$ItableFor((tinfo*)&ityp, (tinfo*)&etyp) \
)

#define ICONVERTIE(iexpr) ({       \
	interface e = iexpr;           \
	(interface){e.val$, TINFO(e)}; \
})

#define ICONVERTEI(iexpr, ityp) ({                          \
	interface e = iexpr;                                    \
	(interface){                                            \
		e.val$,                                             \
		builtin$ItableFor((tinfo*)&ityp, (tinfo*)(e.itab$)) \
	};                                                      \
})


#define ICONVERTII(iexpr, ityp) ({                         \
	interface e = iexpr;                                   \
	(interface){                                           \
		e.val$,                                            \
		builtin$ItableFor((tinfo*)&ityp, (tinfo*)TINFO(e)) \
	};                                                     \
})

static inline
bool implements(const builtin$Type* t, const builtin$Type * it) {
	return builtin$Type$Implements((builtin$Type*)t, (builtin$Type*)it);
}