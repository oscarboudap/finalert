import joblib
from fastapi import FastAPI
from pydantic import BaseModel
import numpy as np

# Crear app FastAPI
app = FastAPI()

# Cargar modelo
model = joblib.load("model_gb.pkl")

# Definir input
class Features(BaseModel):
    delta_1: float
    delta_5: float
    delta_10: float
    delta_30: float

# Endpoint de predicción
@app.post("/predict")
async def predict(features: Features):
    input_array = np.array([[features.delta_1, features.delta_5, features.delta_10, features.delta_30]])
    prediction = model.predict(input_array)
    return {
        "predicted_change": prediction[0]
    }
