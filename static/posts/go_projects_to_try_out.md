# **Easy to Hard GO projects to level UP your game**

Building **projects** is one of the most effective and rewarding ways
to learn new **skills** at your own pace. Plus, if you're on the hunt
for a job, they make a great addition to your application and can
genuinely help you **stand out**.

I've put together a curated list of **five projects** that should
be fun to build and challenging enough to help you level up without
being frustratingly hard.

Before you begin, here are a few **recommendations** to set you up for success:

---

## **1. DO *NOT* start coding right away**

Take a piece of paper and start sketching out what you think you’ll need to
build to complete the project. In this phase, it’s *fundamental* to break
things down into **small, incremental tasks** that combine into a finished product.

*Example:*  
If you want to build a blog, you might create tasks like
these (this is what professionals do too):

1. Create Git repository.  
2. Initialize the Go project with a simple “Hello, world.”  
3. Create a basic Go server that listens on `:8080`.  
4. Implement home page handler.  
5. Implement post page handler.  
6. Add persistent storage (SQLite, Postgres, JSON file, Redis).  
7. Implement login functionality.  
8. Add security by encrypting user passwords.  
9. Add JWT authentication.

---

## **2. DO *NOT* use AI assistance**

No coding agents, no chat based helpers, no autocomplete.

AI tends to complicate things, causing the project to grow faster than your understanding.
Before long, you won’t be able to modify your own code simply because you
didn’t build the foundations yourself.

---

## **3. Keep things simple**

Don’t overcomplicate anything in the beginning. Start with **small, digestible chunks**.
If you overthink or overengineer every step, you’ll quickly get discouraged.
The project’s complexity should grow together with your skills.
If you don’t know how to build a simple web server in Go, don’t start by attempting
one with 25 handlers, 10 middlewares, and all the bells and whistles that make it
production ready. Get a **basic “Hello, world” page** working first.
Then change things **one step at a time**.

---

## **4. Stick to the standard library**

You’re learning. Using libraries built by others won’t help you grow
your skills right now. The Go standard library is a **piece of art** and contains
everything you need for most beginner and intermediate projects. At this stage,
it’s better to write a *crappy, insecure password encryptor yourself* than to
use a perfect one that teaches you nothing about how encryption actually works.

---

## **1. [EASY] Replicate the `cat` Utility from Linux**

As a beginner, you may not yet have the skills to design a tool from scratch.  
So instead of freezing before you even begin, **copy something that already exists**
and is easy to understand. The `cat` utility is perfect for this.

---

## **2. [EASY] URL Shortener**

The classic system design interview question.

**Concept:** Map a short code (e.g., `xh5a`) to a long URL.

```go
type URLMap struct {
    ShortCode string `json:"short_code"`
    Original  string `json:"original"`
}

