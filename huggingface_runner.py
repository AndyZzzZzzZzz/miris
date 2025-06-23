from transformer import AutoTokenizer, AutoModelForCausalLM
import torch
import sys
import json

# other candidate: meta-llama/Llama-2-7b-hf
# mistralai/istral-7B-v0.1
model_name = "tiiuae/falcon-7b"

# loads the tokenizer associated with the model (used to convert text to token IDs)
tokenizer = AutoTokenizer.from_pretrained(model_name)

# loads the actual causal language model (for next-token prediction)
model = AutoModelForCausalLm.from_pretrained(
        model_name,
        torch_dtype=torch.float16,          # loads weights in half-precision to reduce GPU memory usuage
        device_map="auto",                  # spreads the model across available GPUs
        trust_remote_code=True              # allows loading models with custom Python code from the model repo
        )


def emb(text):
    ids = tok(text, return_tensors="pt").to(net.device)

    with torch.no_grad():
        out = net(**ids).last_hidden_state.mean(1)
    return out.squeeze().cpu().half().numpy().tolist()


def generate_dummy_embedding(input_text: str):
    """ 
        Simulates the kind of output a GNN or visual encoder would return for a given object or node

        1. Tokenization: Converts input text into token IDs
        2. Model inference: 
            Runs a forward pass through the model
            with torch.no_grad() disable gradient computation (saves memory)
        3. Extract output embedding:
            output.last_hidden_state is the output embedding for each token
            mean(dim=1) averages token embeddings to get a single vector 
            squeeze() removes extra dimensions
            cpu().numpy() converts the result to a NumPy array
    """

    input_ids = tokenizer(input_text, return_tensor="pt").input_ids.to(model.device)
    with torch.no_grad():
        output = model(input_ids)
    return output.last_hidden_state.mean(dim=1).squeeze().cpu().numpy()

if __name__ == "__main__":
    # text = sys.argv[1] if len(sys.argv) > 1 else "hello"
    # embedding = generate_dummy_embedding(text)
    # print("Embedding shape:", embeddig.shape)
    print(json.dumps(emb(sys.stdin.read().strip())))

