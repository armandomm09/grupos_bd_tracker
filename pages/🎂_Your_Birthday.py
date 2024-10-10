import streamlit as st
import streamlit_phone_number as st_phone
import requests
import datetime
import time
import json
st.set_page_config(page_title="Set your BD", page_icon="🎂")

st.title("Let grupos know you...")

st.text("")

######### NAME
st.markdown("### What is your name?")
name = st.text_input("Complete name")

######### EMAIL
st.markdown("### Email")
email = st.text_input("So we can remind you of you partner's birthday 😙")

######### CELLPHONE

st.markdown("### Cellphone")
cellphone = st_phone.st_phone_number("Phone", default_country="MX")
# cellphone = st.number_input("Just to make sure your BFF doesn't hates you cause you forgot 🙄",  help="Just the city's lada. E.g. 2211745452 ")

######### INSTAGRAM

st.markdown("### Your instagram username")
instagram = st.text_input("So every one can post yoy a birthday story", help="Without '@'. E.g. armando_mm09")


######### INFO
st.markdown("### About you")
info = st.text_area("Fun fact or smth", help="")

######### BIRTHDAY
st.markdown("### And your birthdate")
birthday = st.date_input("Duhh 👶", min_value=datetime.date(1950, 1, 1))

st.text("")
st.text("")
st.text("")

def sendForm():
    
    url = 'http://10.50.84.212:4000/api/person'
    person = {
    "name": str(name),
    "birthday": int(2345678), # int(time.mktime(birthday.timetuple())),
    "email": str(email),
    "insta": str(instagram),
    "telephone": str(cellphone["number"]),
    "info": str(info)
    }
    try:
        x = requests.post(url, json=person)
    except requests.exceptions.RequestException as e:
        print(e)

    print(x.status_code)

st.button("Submit", on_click=sendForm)

