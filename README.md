<picture>
 <source media="(prefers-color-scheme: dark)" srcset="assets/spiderboss_logo_dark.svg">
 <source media="(prefers-color-scheme: light)" srcset="assets/spiderboss_logo.svg">
 <img alt="Spiderboss Logo" src="assets/spiderboss_logo.svg">
</picture>

Spiderboss Lib
==============

## Overview
Spiderboss is a library for controlling cable-suspended platform. Think of the
NFL skycam, that is a cable-suspended camera.

As of now there aren't any specific target applications and this is just for-fun
project. This is still a WIP (work in progress).

## Concepts
The main parts are the bound box, the virtually modeled real-world measurements
of the cable system. The winches are a combination of a stepper motor and a
spool. Spools can be of arbitrary size, with arbitrary sized cable. The spool
mechanism will:
- calculate maximum length of cable the spool can hold
- adjust the revolutions required to extend a certain amount of cable based on
  the current diameter of the spool.

The main functionality of this library is to return the number of steps each
motor need to take to move the carrier in an Cartesian coordinate bound box and
the delay each motor needs to have in order to simultaneously finish their
steps.

## Architecture
```mermaid
---
title: Spiderboss Class Diagram
---

classDiagram
  LocVec ..> Units
  Spool ..> Units
  Winch ..> Spool
  Winch ..> Motor
  Winch ..> Movement
  BoundBox ..> Units
  BoundBox ..> LocVec
  Carrier ..> Winch
  Carrier ..> Movement
  Carrier ..> BoundBox
  Carrier ..> Serial

  class Serial{
    <<interface>>
    io.ReadWriter
  }
  class Spool{
    +Basis EmptyDiam
    +Basis FullDiam
    +Basis EmptyLength
    +Basis FullLength
    +Basis CableDiam
    +Basis InitCableOneSpool
    +Revolutions(Basis) float64, error
    +RemoveCable(Basis) error
    +AddCable(Basis) error
    +CableOnSpool() Basis
  }


  class Motor{
    +int64 StepsPerRev
    +time.Duration MinStepDelay
    +FastestDuration(float64) time.Duration
    +Steps(float64) int64
  }

  class Carrier{
    +List~*Winch~ Winches
    +BoundBox BoundBox
    +Serial Serial
    +*log.Logger Logger
    +LineTo()
  }

  class Units {
    +int64 Basis
    +Basis Millimeter
    +Basis Centimeter
    +Basis Meter
    +Basis Inch
    +Basis Foot
    +Basis Yard
  }

  class LocVec {
    +Basis X
    +Basis Y
    +Basis Z
    +Len() Basis
    +AbsSubtract(LocVec) *LocVec
    +GreaterThan(LocVec) bool
  }

  class Winch {
    +Motor Motor
    +Spool Spool
    +Basis Origin
    +Move(Basis, time.Duration) *Movement, error
    +Extend(Basis, time.Duration) *Movement, error
    +Retract(Basis, time.Duration) *Movement, error
    +SetLength(Basis, time.Duration) *Movement, error
  }

  class Movement{
    +int64 Steps
    +time.Duration Delay
    +itoa Direction
    +Bytes() []byte
  }

  class BoundBox{
    +List~LocVec~ Origins
    +CableLenHome() Basis
    +CableLenAt(LocVec) Basis, error
  }
```

## License
Apache License 2.0
