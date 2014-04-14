# Go [Artemis] - Spaceship warriors #

Port of the Sapceship warriors Artemis example.
It uses the Go port of Artemis, and implements almost everything of the original example - except the HUD and in general text stuff, because that's boring.

The only notable change, is that there is a different particle render system - implemented mostly as an exercise.

## Todo ##

Revisit rendering system.

- Stars should be renderered as particles.
- Add pool for sf.Vertex2f.
- ParticleRenderer should just create vertexarrays.
- SpriteRenderer should just create sprites.
- A new system (SfmlRenderer) should get sprites, arrays
and render them in layer order.
