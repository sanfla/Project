import pandas as pd
import numpy as np
import nltk
import re
import streamlit as st
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.naive_bayes import MultinomialNB
from nltk.tokenize import word_tokenize
from nltk.corpus import stopwords

# Preprocessing Teks
stop_words = set(stopwords.words('english'))

def preprocess_text(text):
    words = word_tokenize(text)
    filtered_words = [word.lower() for word in words if word.isalpha() and word.lower() not in stop_words]
    return ' '.join(filtered_words)

# Fungsi untuk memuat dan memproses data
def load_and_preprocess_data(train_url, validation_url):
    df_train = pd.read_csv(train_url)
    df_validation = pd.read_csv(validation_url)

    df_train['clean_text'] = df_train['text'].apply(preprocess_text)
    df_validation['clean_text'] = df_validation['text'].apply(preprocess_text)

    X_train, y_train = df_train['clean_text'], df_train['sentiment']
    X_test, y_test = df_validation['clean_text'], df_validation['sentiment']

    return X_train, X_test, y_train, y_test

# Fungsi untuk melatih model
def train_model(X_train, y_train):
    tfidf_vectorizer = TfidfVectorizer(max_features=5000)
    X_train_tfidf = tfidf_vectorizer.fit_transform(X_train)

    nb_model = MultinomialNB()
    nb_model.fit(X_train_tfidf, y_train)

    return tfidf_vectorizer, nb_model

# Fungsi untuk prediksi sentimen dengan empat kelas
def predict_sentiment(text, tfidf_vectorizer, nb_model):
    cleaned_text = preprocess_text(text)
    vectorized_text = tfidf_vectorizer.transform([cleaned_text])
    prediction = nb_model.predict(vectorized_text)[0]

    return prediction

# Fungsi untuk mendapatkan respons acak berdasarkan sentimen
def get_random_response(sentiment, df_response):
    relevant_responses = df_response[df_response['sentiment'] == sentiment]['response']
    if not relevant_responses.empty:
        return np.random.choice(relevant_responses)
    else:
        return "Maaf, kami belum dapat merekomendasikan kepada anda."

# Load and preprocess data for the first feature (Rekomendasi Makanan)
train_url_makanan = "https://raw.githubusercontent.com/SyahrezaAdnanAlAzhar/tugasBesarLabAI/main/makanan_sehat_training.csv"
validation_url_makanan = "https://raw.githubusercontent.com/SyahrezaAdnanAlAzhar/tugasBesarLabAI/main/makanan_sehat_validation.csv"
X_train_makanan, X_test_makanan, y_train_makanan, y_test_makanan = load_and_preprocess_data(train_url_makanan, validation_url_makanan)
tfidf_vectorizer_makanan, nb_model_makanan = train_model(X_train_makanan, y_train_makanan)

# Load and preprocess data for the second feature (Rekomendasi Olahraga)
train_url_olahraga = "https://raw.githubusercontent.com/sanfla/Project/main/TugasAILab/olahraga_training%20-%20Sheet1.csv"
validation_url_olahraga = "https://raw.githubusercontent.com/sanfla/Project/main/TugasAILab/olahraga_validation%20-%20Sheet1.csv"
X_train_olahraga, X_test_olahraga, y_train_olahraga, y_test_olahraga = load_and_preprocess_data(train_url_olahraga, validation_url_olahraga)
tfidf_vectorizer_olahraga, nb_model_olahraga = train_model(X_train_olahraga, y_train_olahraga)

# Load and preprocess data for the third feature (Larangan Terhadap Penyakit)
train_url_larangan = "https://raw.githubusercontent.com/adeikmalm/sedangsakitnew/main/sedangsakit_training.csv"
validation_url_larangan = "https://raw.githubusercontent.com/SyahrezaAdnanAlAzhar/tugasBesarLabAI/main/sedangsakit_validation.csv"
X_train_larangan, X_test_larangan, y_train_larangan, y_test_larangan = load_and_preprocess_data(train_url_larangan, validation_url_larangan)
tfidf_vectorizer_larangan, nb_model_larangan = train_model(X_train_larangan, y_train_larangan)

# Load response data from all features
df_makanan_response = pd.read_excel("https://docs.google.com/spreadsheets/d/1WIHxAdAIhbcamQH0bBQ8eF6Pv2OD32tvS9bnAFqtgC4/export?format=xlsx")
df_olahraga_response = pd.read_csv("https://raw.githubusercontent.com/sanfla/Project/main/TugasAILab/response_olahraga%20-%20Sheet1%20(1).csv")
df_larangan_response = pd.read_excel("https://docs.google.com/spreadsheets/d/1BWlFGTsygh5mAemIeZKIGUG1THbt-5d5pE5XB6rtv-w/export?format=xlsx")

# Streamlit app
st.title("NutriNest")

# Inisialisasi session state
if "messages" not in st.session_state:
    st.session_state.messages = []

# Display chat messages from history on app rerun
for message in st.session_state.messages:
    with st.chat_message(message["role"]):
        st.markdown(message["content"])

# Accept user input
if prompt := st.chat_input("Halo aku asisten kesehatan ada yang bisa saya bantu?"):
    # Add user message to chat history
    st.session_state.messages.append({"role": "user", "content": prompt})
    
    # Determine which dataset to use based on user input
    if "makanan" in prompt.lower():
        X_train, X_test, y_train, y_test = X_train_makanan, X_test_makanan, y_train_makanan, y_test_makanan
        tfidf_vectorizer, nb_model = tfidf_vectorizer_makanan, nb_model_makanan
        df_response = df_makanan_response
    elif "olahraga" in prompt.lower() or "aktivitas" in prompt.lower():
        X_train, X_test, y_train, y_test = X_train_olahraga, X_test_olahraga, y_train_olahraga, y_test_olahraga
        tfidf_vectorizer, nb_model = tfidf_vectorizer_olahraga, nb_model_olahraga
        df_response = df_olahraga_response
    elif "larangan" in prompt.lower() or "sakit" in prompt.lower():
        X_train, X_test, y_train, y_test = X_train_larangan, X_test_larangan, y_train_larangan, y_test_larangan
        tfidf_vectorizer, nb_model = tfidf_vectorizer_larangan, nb_model_larangan
        df_response = df_larangan_response
    else:
        # Use default dataset (for example, makanan dataset)
        X_train, X_test, y_train, y_test = X_train_makanan, X_test_makanan, y_train_makanan, y_test_makanan
        tfidf_vectorizer, nb_model = tfidf_vectorizer_makanan, nb_model_makanan
        df_response = df_makanan_response

    prediction_output = predict_sentiment(prompt, tfidf_vectorizer, nb_model)
    random_response = get_random_response(prediction_output, df_response)

    # Display user message in chat message container
    with st.chat_message("user"):
        st.markdown(prompt)
    # Add assistant message to chat history
    st.session_state.messages.append({"role": "assistant", "content": random_response})
    
    # Display assistant response in chat message container
    with st.chat_message("assistant"):
        st.markdown(random_response)