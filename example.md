## ADK Features

Based on the documentation, ADK delivers a comprehensive set of features for building, evaluating, and deploying sophisticated AI agents. Here is a summary of the key capabilities:

### Core Agent Architecture
*   **LLM Agents**: The primary "thinking" unit of ADK, powered by a Large Language Model for reasoning, decision-making, and tool use.
*   **Multi-Agent Systems**: Build complex applications by composing multiple, specialized agents in a hierarchy, enabling delegation and collaboration.
*   **Workflow Agents**: Orchestrate the execution flow of other agents using deterministic patterns:
    *   **Sequential Agents**: Run sub-agents one after another.
    *   **Parallel Agents**: Run sub-agents concurrently.
    *   **Loop Agents**: Run sub-agents in an iterative loop.
*   **Custom Agents**: Implement unique, non-LLM based orchestration logic by extending the `BaseAgent` class for ultimate flexibility.
*   **Agent Config (YAML)**: Build agents declaratively using YAML configuration files, minimizing the need for boilerplate code.

### Tooling and Integrations
*   **Function Tools**: Extend agent capabilities by creating custom tools from standard functions.
*   **Long-Running Function Tools**: Handle asynchronous tasks that take significant time to complete without blocking the agent.
*   **Agents-as-a-Tool**: Use an entire agent as a callable tool within another agent.
*   **Built-in Tools**: Leverage pre-built capabilities like Google Search, Vertex AI Search, and secure code execution.
*   **Google Cloud Tools**: Seamlessly connect to Google Cloud services like API Hub, Application Integration, and databases via MCP Toolbox.
*   **Third-Party Tools**: Integrate tools from popular frameworks like LangChain and CrewAI.
*   **MCP (Model Context Protocol) Tools**: Use and expose tools via the open MCP standard.
*   **Tool Confirmation**: Implement human-in-the-loop workflows by requiring user confirmation before a tool executes.
*   **Tool Authentication**: Securely manage credentials for tools that access protected APIs.

### Runtime and Context Management
*   **Session & State**: Manage conversational context with a robust session service that tracks history (`Events`) and maintains a working memory (`State`) for each interaction.
*   **Memory**: Enable agents to recall information across multiple sessions using a long-term, searchable knowledge store.
*   **Artifacts**: Allow agents to save, load, and manage versioned files and binary data (e.g., images, PDFs) associated with a session or user.
*   **Callbacks & Plugins**: Hook into the agent's execution lifecycle to implement logging, guardrails, caching, and other cross-cutting concerns.

### Advanced Capabilities
*   **Streaming (Live/Bidi)**: Build real-time, interactive experiences with native support for bidirectional streaming of text, audio, and video.
*   **Streaming Tools**: Create tools that can stream intermediate results back to the agent for real-time monitoring and reaction.
*   **Multi-Model Support**: Flexibly use different LLMs (Gemini, GPT, Claude, etc.) within your agents, powered by integrations like LiteLLM.

### Development Lifecycle
*   **Evaluation**: Systematically assess agent performance by evaluating both the final response quality and the step-by-step execution trajectory.
*   **Observability & Logging**: Gain deep insights into agent behavior with detailed logging and integrations with observability platforms like AgentOps, Arize AX, Cloud Trace, and Weave.
*   **Deployment**: Deploy your agents to a variety of environments, including serverless (Cloud Run), Kubernetes (GKE), and the fully managed Vertex AI Agent Engine.
*   **Agent2Agent (A2A) Protocol**: Build distributed, multi-agent systems where agents running as separate services can communicate and collaborate using the open A2A standard.

## Feature Support Matrix

The following matrix shows which ADK features are currently supported in each language runtime. Support is determined by the availability of code samples in the official documentation.

| Feature Category | Feature | Python (`adk-python`) | Java (`adk-java`) |
| :--- | :--- | :--- | :--- |
| **Core Agent Architecture** | **LLM Agents** | [Supported](https://google.github.io/adk-docs/agents/llm-agents/#defining-the-agents-identity-and-purpose) | [Supported](https://google.github.io/adk-docs/agents/llm-agents/#defining-the-agents-identity-and-purpose) |
| | **Multi-Agent Systems** | [Supported](https://google.github.io/adk-docs/agents/multi-agents/#11-agent-hierarchy-parent-agent-sub-agents) | [Supported](https://google.github.io/adk-docs/agents/multi-agents/#11-agent-hierarchy-parent-agent-sub-agents) |
| | **Workflow: Sequential Agents** | [Supported](https://google.github.io/adk-docs/agents/workflow-agents/sequential-agents/#full-example-code-development-pipeline) | [Supported](https://google.github.io/adk-docs/agents/workflow-agents/sequential-agents/#full-example-code-development-pipeline) |
| | **Workflow: Loop Agents** | [Supported](https://google.github.io/adk-docs/agents/workflow-agents/loop-agents/#full-example-iterative-document-improvement) | [Supported](https://google.github.io/adk-docs/agents/workflow-agents/loop-agents/#full-example-iterative-document-improvement) |
| | **Workflow: Parallel Agents** | [Supported](https://google.github.io/adk-docs/agents/workflow-agents/parallel-agents/#full-example-parallel-web-research) | [Supported](https://google.github.io/adk-docs/agents/workflow-agents/parallel-agents/#full-example-parallel-web-research) |
| | **Custom Agents** | [Supported](https://google.github.io/adk-docs/agents/custom-agents/#part-1-simplified-custom-agent-initialization) | [Supported](https://google.github.io/adk-docs/agents/custom-agents/#part-1-simplified-custom-agent-initialization) |
| | **Agent Config (YAML)** | [Supported](https://google.github.io/adk-docs/agents/config/#build-an-agent) | [No](https://google.github.io/adk-docs/agents/config/#known-limitations) |
| **Tooling & Integrations** | **Function Tools** | [Supported](https://google.github.io/adk-docs/tools/function-tools/#example) | [Supported](https://google.github.io/adk-docs/tools/function-tools/#example) |
| | **Long-Running Function Tools** | [Supported](https://google.github.io/adk-docs/tools/function-tools/#creating-the-tool) | [Supported](https://google.github.io/adk-docs/tools/function-tools/#creating-the-tool) |
| | **Agents-as-a-Tool** | [Supported](https://google.github.io/adk-docs/tools/function-tools/#example) | [Supported](https://google.github.io/adk-docs/tools/function-tools/#example) |
| | **Built-in: Google Search** | [Supported](https://google.github.io/adk-docs/tools/built-in-tools/#google-search) | [Supported](https://google.github.io/adk-docs/tools/built-in-tools/#google-search) |
| | **Built-in: Code Execution** | [Supported](https://google.github.io/adk-docs/tools/built-in-tools/#code-execution) | [Supported](https://google.github.io/adk-docs/tools/built-in-tools/#code-execution) |
| | **Built-in: GKE Code Executor** | [Supported](https://google.github.io/adk-docs/tools/built-in-tools/#gke-code-executor) | Planned |
| | **Built-in: Vertex AI RAG** | [Supported](https://google.github.io/adk-docs/tools/built-in-tools/#vertex-ai-rag-engine) | Planned |
| | **Built-in: Vertex AI Search** | [Supported](https://google.github.io/adk-docs/tools/built-in-tools/#vertex-ai-search) | Planned |
| | **Built-in: BigQuery, Spanner, Bigtable** | [Supported](https://google.github.io/adk-docs/tools/built-in-tools/#bigquery) | Planned |
| | **Google Cloud Tools (API Hub, etc.)** | [Supported](https://google.github.io/adk-docs/tools/google-cloud-tools/#apigee-api-hub-tools) | Planned |
| | **Third-Party Tools (LangChain, CrewAI)** | [Supported](https://google.github.io/adk-docs/tools/third-party-tools/#1-using-langchain-tools) | Planned |
| | **MCP Tools** | [Supported](https://google.github.io/adk-docs/tools/mcp-tools/#example-1-file-system-mcp-server) | [Supported](https://google.github.io/adk-docs/tools/mcp-tools/#example-1-file-system-mcp-server) |
| | **Tool Confirmation** | [Supported](https://google.github.io/adk-docs/tools/confirmation/#boolean-confirmation) | Planned |
| | **Tool Authentication** | [Supported](https://google.github.io/adk-docs/tools/authentication/#1-configuring-tools-with-authentication) | Planned |
| **Runtime & Context** | **Session & State Management** | [Supported](https://google.github.io/adk-docs/sessions/session/#example-examining-session-properties) | [Supported](https://google.github.io/adk-docs/sessions/session/#example-examining-session-properties) |
| | **Memory** | [Supported](https://google.github.io/adk-docs/sessions/memory/#example-adding-and-searching-memory) | Planned |
| | **Artifacts** | [Supported](https://google.github.io/adk-docs/artifacts/#what-are-artifacts) | [Supported](https://google.github.io/adk-docs/artifacts/#what-are-artifacts) |
| | **Callbacks** | [Supported](https://google.github.io/adk-docs/callbacks/types-of-callbacks/#before-agent-callback) | [Supported](https://google.github.io/adk-docs/callbacks/types-of-callbacks/#before-agent-callback) |
| | **Plugins** | [Supported](https://google.github.io/adk-docs/plugins/#define-and-register-plugins) | Planned |
| **Advanced Capabilities** | **Streaming (Live/Bidi)** | [Supported](https://google.github.io/adk-docs/get-started/streaming/quickstart-streaming/#2-project-structure) | [Supported](https://google.github.io/adk-docs/get-started/streaming/quickstart-streaming-java/#creating-an-agent) |
| | **Streaming Tools** | [Supported](https://google.github.io/adk-docs/streaming/streaming-tools/) | Planned |
| | **Multi-Model: LiteLLM** | [Supported](https://google.github.io/adk-docs/agents/models/#using-cloud-proprietary-models-via-litellm) | Planned |
| | **Multi-Model: Direct Anthropic** | Planned | [Supported](https://google.github.io/adk-docs/agents/models/#using-anthropic-models) |
| **Development Lifecycle** | **Evaluation** | [Supported](https://google.github.io/adk-docs/evaluate/#how-evaluation-works-with-the-adk) | Planned |
| | **Observability (AgentOps, Arize, etc.)** | [Supported](https://google.github.io/adk-docs/observability/agentops/) | Planned |
| | **Deployment: Vertex AI Agent Engine** | [Supported](https://google.github.io/adk-docs/deploy/agent-engine/#standard-deployment) | [No](https://google.github.io/adk-docs/deploy/agent-engine/) |
| | **Deployment: Cloud Run** | [Supported](https://google.github.io/adk-docs/deploy/cloud-run/#deployment-commands) | [Supported](https://google.github.io/adk-docs/deploy/cloud-run/#deployment-commands) |
| | **Deployment: GKE** | [Supported](https://google.github.io/adk-docs/deploy/gke/#option-1-manual-deployment-using-gcloud-and-kubectl) | Planned |
| | **Agent2Agent (A2A) Protocol** | [Supported](https://google.github.io/adk-docs/a2a/quickstart-consuming/#getting-the-sample-code) | Planned |