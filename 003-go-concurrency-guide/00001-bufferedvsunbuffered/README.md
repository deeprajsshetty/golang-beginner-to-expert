# Golang Concurrency: Buffered vs Unbuffered Channels

This sample program demonstrates how **Buffered** and **Unbuffered** channels work in Golang, including the use of **WaitGroup** to manage goroutine completion in a concurrent setting.

---

## 🧠 Key Concepts Covered

- **Buffered Channels**: Allow sending data without a corresponding receiver (up to buffer capacity).
- **Unbuffered Channels**: Require a receiver to be ready before sending proceeds.
- **WaitGroup**: Helps wait for a collection of goroutines to finish executing.

---

# 🧵 Golang Concurrency: Buffered vs Unbuffered Channels

This code sample demonstrates how Golang handles concurrency using **Buffered** and **Unbuffered** Channels and highlights the importance of **`sync.WaitGroup`**.

---

## 🔍 Function Breakdown

### ✅ Buffered Channel Examples (with WaitGroup)

- `WaitGroup` is used to **block the main goroutine** until all fetch routines finish.
- `respch` is **shared among all goroutines**, so all results go into a single channel.
- **Buffered channels** allow goroutines to send messages **even if the main goroutine hasn't started reading**, up to the buffer's capacity.

---

### ✅ Unbuffered Channel Examples

- Each routine uses its **own dedicated unbuffered channel**.
- The **main goroutine must be ready to receive**, or the goroutine will **block until the receiver is available**.

---

## 🧩 Why Use `WaitGroup` with Buffered Channels?

Even though **Buffered Channels** do not block until the buffer is full, using a **WaitGroup** ensures:

- ✅ All goroutines **complete their execution**, even if the channel has enough buffer space.
- ✅ It ensures **data consistency** and avoids **premature channel closing**, which would otherwise cause a **panic** if goroutines are still writing.

> ⚠️ Without `WaitGroup`, we cannot guarantee that all goroutines have finished writing to the buffered channel before it is closed.

---

## ✅ Best Practices

- ✅ Use **`WaitGroup`** to ensure all goroutines complete before continuing or closing channels.
- ✅ Always **close buffered channels** if you're using a `for-range` loop to read from them.
- ✅ Prefer **Unbuffered Channels** when you need **tight synchronization** between sender and receiver.
- ✅ Use **separate unbuffered channels** for better isolation and to prevent blocking across goroutines.

---

> This document serves as a handy reference for Golang concurrency and channel behavior, especially for comparing buffered vs. unbuffered communication models.
