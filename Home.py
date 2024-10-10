import streamlit as st
import streamlit_calendar as st_calendar
import time 

st.set_page_config(page_title="Calendar" , page_icon="📆", )

st.markdown("Welcome to... ")
st.markdown("# Grupos estudiantiles Birthday tracker")
st.text("")
st.text("")
st.text("")
st.text("")
st_calendar.calendar(
    events=[],
    options={},
    )

