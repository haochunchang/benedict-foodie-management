import axios from "axios";
import { useState } from "react";
import { Calendar } from "react-native-calendars";

const FoodCalendar = ({ createForm, backendUrl }) => {

    // Fetch current month record
    const currentYear = new Date().getFullYear();
    const currentMonth = new Date().getMonth() + 1;
    let currentRecords = {};
    axios.get(`${backendUrl}/records/${currentYear}/${currentMonth}`)
        .then(response => response.json())
        .catch(err => alert(err))
        .then(result => {
            for (const res of result) {
                currentRecords[res.eating_date] = res;
                currentRecords[res.eating_date].isModifying = true;
            }
        });

    let initMarkedDates = {};
    for (r in currentRecords) {
        initMarkedDates[r] = { marked: true, selectedColor: 'blue' }
    }
    const [markedDates, onChangeMarkedDates] = useState(initMarkedDates);

    return <Calendar
        onDayPress={({ dateString }) => {
            if (dateString in markedDates) {
                createForm(currentRecords[dateString])
            } else {
                createForm({
                    Name: "",
                    EatingDate: dateString,
                    EatenQuantity: 0,
                    SatisfactionScore: 0,
                    isModifying: false
                });
            }
        }}
        markedDates={markedDates}
        // Month format in calendar title. Formatting values: http://arshaw.com/xdate/#Formatting
        monthFormat={'yyyy-MMM'}
        hideExtraDays={true}
        firstDay={1}
        hideDayNames={false}
        onPressArrowLeft={subtractMonth => subtractMonth()}
        onPressArrowRight={addMonth => addMonth()}
        disableAllTouchEventsForDisabledDays={true}
        enableSwipeMonths={true}
    />
}

export default FoodCalendar