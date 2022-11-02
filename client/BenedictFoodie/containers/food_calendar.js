import axios from "axios";
import { useEffect, useState } from "react";
import { View, Modal } from 'react-native';
import { Calendar } from "react-native-calendars";
import { RecordForm } from './forms';


const FoodCalendar = ({ backendUrl }) => {

    const defaultRecord = {
        Name: "",
        EatingDate: "",
        EatenQuantity: 0,
        SatisfactionScore: 0,
        isModifying: false
    };
    const [curRecord, onChangeCurRecord] = useState(defaultRecord);
    const [thisMonthRecord, onChangeThisMonthRecord] = useState({});
    const [markedDates, onChangeMarkedDates] = useState({});

    useEffect(() => {
        fetchCurrentMonthRecord(backendUrl, onChangeThisMonthRecord);
    }, []);
    useEffect(() => {
        onChangeMarkedDates(getMarkedDates(thisMonthRecord));
    }, [thisMonthRecord, curRecord]);

    return (
        <View>
            <Modal visible={curRecord.EatingDate !== ""}>
                <RecordForm
                    record={curRecord}
                    closeForm={() => { onChangeCurRecord(defaultRecord) }}
                    backendUrl={backendUrl}
                    thisMonthRecord={thisMonthRecord}
                    onChangeThisMonthRecord={onChangeThisMonthRecord}
                />
            </Modal>
            <Calendar
                onDayPress={({ dateString }) => {
                    if (dateString in markedDates) {
                        onChangeCurRecord(thisMonthRecord[dateString])
                    } else {
                        onChangeCurRecord({
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
        </View>
    )
}

const fetchCurrentMonthRecord = (backendUrl, onChangeThisMonthRecord) => {
    const currentYear = new Date().getFullYear();
    const currentMonth = new Date().getMonth() + 1;
    let currentRecords = {};
    fetch(`${backendUrl}/records/${currentYear}/${currentMonth}`, {
        method: 'GET',
        headers: {
            Accept: 'application/json',
            'Content-Type': 'application/json'
        }
    })
        .then(response => response.json())
        .then(result => {
            for (const res of result) {
                currentRecords[res.EatingDate] = res;
                currentRecords[res.EatingDate].isModifying = true;
            }
            onChangeThisMonthRecord(currentRecords);
        }).catch(err => alert(err));
}


const getMarkedDates = (records) => {
    let marked = {};
    for (r in records) {
        marked[r] = { marked: true, selectedColor: 'blue' };
    }
    return marked;
}

export default FoodCalendar