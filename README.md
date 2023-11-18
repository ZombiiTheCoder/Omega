# Omega Language

## Rule \#1 All Functions Must Return Something
```
fn Hello_World() -> void {
    return null;
}
```

## Rule \#2 All Expression and or Statements Must End With A Semicolon
```
fn Hello_World() -> void {
    int i = 0 // Invalid
    int i = 0; // Valid
    return null // Invalid
    return null; // Valid
} // Invalid
}; // Valid
```

## Functions (Types)
```

// fn name(args) {} is the same as fn name(args) -> void {}

fn Hello_World() -> void {
    int i = 0;
    return null;
};

fn Hello_World2() {
    int i = 0;
    return null;
};

// Both Are Valid Functions
```

## Types
```
// the basic types that exist are

Void       // Null
Int        // Integer64
Float      // Float64
Func       // Functions
NativeFunc // Native Function
Str        // String
Bool       // Boolean

```