# Philosophy

## Goal and rationale

**TL;DR:** Make an easily expandable collection of simple reusable FFmpeg recipes.

I often need to do specific jobs, repeatedly, and specific to a temporary scenario. E.g. horizontally stack a bunch of videos from differernt folders, write specific metadata, add covers, convert image sequences to videos at specific framerates, ...

For years I wrote tons of little scripts such as in [Github/Lauloque/FFmpeg-bat-collection](https://github.com/Lauloque/FFmpeg-bat-collection), which are good for archival and showing the "FFmpeg recipe" I've found for each need. But they are terrible to reuse, and a pain to deploy as I work on multiple machines with different OSes and environments.

The best recipes are often found in a family cookbook: handrwitten, currated over generations, and refined every time they're used.

**FFrecipes** follows the same idea. It's both a collection of reusable FFmpeg recipes, and the cookbook that lets you see all the available recipes, what they do, how they can be used, lets you cook with them, and you can easily add new ones.

## Non-goal

This won't be a graphical media encoder. There is already Handbrake, Fastflix, IFME, Xmedia Recode, Adobe ME, ... Nobody needs yet another one, and FFrecipes will not replace those.
