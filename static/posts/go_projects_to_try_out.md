# **How to approach a programming project**

## Crappy code is better than AI code

---

Building **projects** is one of the most effective and rewarding ways
to learn new **skills** at your own pace,
BUT, and that's a big BUTT, only if you do it **right**.
This guide will help you navigate the best way to pick a project,
stick with it, and learn new skills by tackling the challenges ahead of you.

I've interviewed many **aspiring developers** in my professional career
and there is a world of difference between those who have put in the work
to create a pet project and those who think that they will learn how to
program once they are hired and that people will spend 6 months teaching
them how not to suck at coding.

But here is the good news. Building a project can be a rewarding and fun
experience if you follow a few simple guidelines.

## **1. DO NOT start coding right away**

Finished projects are a **complex** result of many days of work. There is no way
you can build a good project without thinking **ahead** about the features
you want to build.

Take a piece of paper and start sketching out what you think you’ll need in
order to complete the project. In this phase, it’s *fundamental* to break
things down into **small, incremental tasks** that combine into a finished product.

If you want to build a blog you don't do that **haphazardly** in a single push,
the same way a furniture maker doesn't assemble and varnish the individual
parts of a cabinet before cutting them to size. If you do that you will get easily
discouraged, hate yourself, your life, and everyone around you.

I'm betting you don't want that. Do you?

## **2. Keep things simple**

Don’t overcomplicate things in the beginning. Start with **small, digestible chunks**.
If you overthink and overengineer every step of the way, you’ll quickly get disheartened.
The project’s complexity should grow together with your skills.

If you don’t know how to build a simple web server in Go, don’t start by attempting
one with 25 handlers, 10 middlewares, and all the bells and whistles that make it
production ready. Get a **basic “Hello, world” page** working first.
Then change things **one step at a time**.

*Example:*

1. Create Git repository.
2. Initialize the Go project with a simple “Hello, world.”
3. Create a basic Go server that listens on `:8080`.
4. Implement home page handler.
5. Implement post page handler.
6. Add persistent storage (SQLite, Postgres, JSON file, Redis).
7. Implement login functionality.
8. Add security by encrypting user passwords.
9. Add JWT authentication.

Tackle one task, and cross it off the list when you are done with it.
Pat yourself on the back, get an ice cream, rinse and repeat.

## **3. DO *NOT* use AI assistance**

No coding agents, no chat based helpers, no autocomplete.

**AI** tends to overcomplicate things, causing the project to grow **faster** than
your understanding. Before long, you won’t be able to modify your own code simply
because you didn’t build the foundations yourself.

Avoid AI in any of its forms like the plague. I know you think AI helps you learn;
it doesn't! And there is no skill involved in writing prompts. The only thing
that matters is knowing what you need. And guess what? You can't know what you
need if you don't learn what you need to know. I'll pause for a moment here
to let that sink in!

...

Good, so **delete** your ChatGPT and Gemini bookmarks, deactivate your Copilot plugin
and close YouTube. And since you are at it, delete all kinds of social media.

## **4. Stick to the standard library**

During the learning phase, using libraries built by others won’t help you grow
your skills. If you program in Go, the Go standard library is a **piece of art**
and contains everything you need for most beginner and intermediate projects.

At this stage, and even if you need stuff that is not in the std library, it’s
better to write a *crappy, insecure password encryptor yourself* than to use a
perfect one that teaches you nothing about how encryption actually works.

**Just remember: never use that crappy version in a real production app,
or you’ll get hacked in seconds. But for a pet project, more than good enough!**

---

Need inspiration for you next project? Take a look at my next article!
