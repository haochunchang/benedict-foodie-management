import { Calendar } from "react-native-calendars";

const FoodCalendar = ({ createForm }) => {
    return <Calendar
        onDayPress={day => {
            createForm(day.dateString);
        }}
        onDayLongPress={day => {
            console.log('selected day', day);
        }}
        // Month format in calendar title. Formatting values: http://arshaw.com/xdate/#Formatting
        monthFormat={'yyyy MM'}
        // Do not show days of other months in month page. Default = false
        hideExtraDays={true}
        // If firstDay=1 week starts from Monday. Note that dayNames and dayNamesShort should still start from Sunday
        firstDay={1}
        hideDayNames={false}
        onPressArrowLeft={subtractMonth => subtractMonth()}
        onPressArrowRight={addMonth => addMonth()}
        disableAllTouchEventsForDisabledDays={true}
        // Replace default month and year title with custom one. the function receive a date as parameter
        renderHeader={date => {
            /*Return JSX*/
        }}
        enableSwipeMonths={true}
    />
}

export default FoodCalendar