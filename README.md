<div align="center">
  <h1>🗣️ MyGO · KVCache-centric Architecture for LLM Serving</h1>
</div>
<br/>

<p align="center">
    <img height="21" src="https://img.shields.io/badge/License-Apache--2.0-ffffff?labelColor=d4eaf7&color=2e6cc4" alt="license">
</p>

<details open>
<summary><b>📕 Table of Contents</b></summary>
  
- 💡 [What is MyGO?](#-what-is-mygo)
- 🎮 [Demo](#-demo)
- 🔎 [Cosine similarity](#-cosine-similarity)
- 🎬 [Get Started](#-get-started)
- 🛠️ [Build from source](#-build-from-source)
- 🛠️ [Launch service from source](#-launch-service-from-source)
- 📚 [Documentation](#-documentation)
- 🏄 [Community](#-community)
- 🙌 [Contributing](#-contributing)
</details>

## 💡 What is MyGO?

MyGO is an LLM API that uses caching and KV-Database to improve response speed, similar to [Gemini](https://ai.google.dev/gemini-api/docs/caching) and [Kimi](https://arxiv.org/pdf/2407.00079v1).

In a typical AI workflow, you might pass the same input Tokens to the model over and over again. With the MyGO API's context caching feature, you can request text multiple times and only need to access the model once. By caching the input Tokens, calculating similarity using [Cosine similarity](#-cosine-similarity), and referencing cached Tokens for subsequent requests, MyGO reduces costs and latency by avoiding the repeated processing of identical input data.

## 🎮 Demo

Check out [Kimi](https://kimi.moonshot.cn/) for a practical example of how MyGO's concepts are applied.

## 🔎 Cosine similarity

Cosine similarity measures the similarity between two vectors by calculating the cosine of the angle between them. The cosine value ranges from -1 to 1, where:

- **1** indicates the vectors are identical in direction,
- **0** indicates the vectors are orthogonal (i.e., at 90°),
- **-1** indicates the vectors are diametrically opposed.

This measure is widely used in positive space, where all values are non-negative, making it particularly useful for comparing textual data in natural language processing tasks.

## 🎬 Get Started

To get started with MyGO, follow these steps:

1. **Install dependencies**: Ensure you have Go installed on your machine.
2. **Clone the repository**:
   ```bash
   git clone https://github.com/Chihaya-Yuka/mygo.git
   cd mygo
   ```
3. **Run the service**:
   ```bash
   go run main.go
   ```

For more detailed instructions, see the [Documentation](#-documentation) section.

## 🛠️ Build from source

If you want to build MyGO from source, follow these steps:

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Chihaya-Yuka/mygo.git
   cd mygo
   ```
2. **Build the project**:
   ```bash
   go build -o mygo main.go
   ```
3. **Run the binary**:
   ```bash
   ./mygo
   ```

This will start the MyGO service on your local machine.

## 🛠️ Launch service from source

To launch the MyGO service from the source, simply run:

```bash
go run main.go
```

This will start the service, which you can then interact with via HTTP requests.

## 📚 Documentation

For more detailed documentation, including API references and advanced configuration options, visit the [MyGO Documentation](https://github.com/Chihaya-Yuka/mygo/wiki).

## 🏄 Community

Join our community to share your experiences, ask questions, and collaborate with others:

- [GitHub Discussions](https://github.com/Chihaya-Yuka/mygo/discussions)

## 🙌 Contributing

We welcome contributions! If you'd like to contribute to MyGO, please:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes and commit them (`git commit -m 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Open a pull request.

