# Questions

## What if user inputs lowercase letters?

## What if user includes invalid commands?

## How exactly does the map of string, interface work?

## Is there a way to have a receiver of slice? Does it matter what type the receiver holds?

    Ex: (cl []Command) func ...
        (sl []string) func ...

## Could reflection be used to map the functions?

## Is it good practice to use the \_ as a way to match function types?

## How can I ignore ascii characters on input?

## Why does draw.Draw take a &Uniform instead of just Uniform

The Uniform struct satisfues the Image interface (ColorModel, Bounds, At) with methods that have a receiver with (c \*Uniform)
(Method sets)[https://go.dev/ref/spec#Method_sets]
