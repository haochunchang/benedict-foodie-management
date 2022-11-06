import React, { useState } from 'react';
import {
    View, Text, TextInput, Button, Alert,
} from "react-native"
import { Select, Option } from '../third_party/react-native-select-list-modified/src';

const SatisfactionScoreDescription = {
    1: "Does not eat with snack added",
    2: "Finish eating with snack",
    3: "Finish eating",
    4: "Eating eagerly",
    5: "Eating eagerly with talking",
}

const HEADERS = {
    Accept: 'application/json',
    'Content-Type': 'application/json'
};

const FoodTypeDropdownList = ({ handle }) => {
    return (
        <Select onSelect={(value) => { handle(value) }}>
            <Option value={'dry'}>Dry Food</Option>
            <Option value={'wet'}>Wet Food</Option>
            <Option value={'snack'}>Snack</Option>
        </Select >
    );
};

const SatisfactionScoreDropdown = ({ initScore, handle }) => {
    return (
        <Select default={SatisfactionScoreDescription[initScore]} onSelect={(value) => { handle(value) }}>
            <Option value={1}>{SatisfactionScoreDescription[1]}</Option>
            <Option value={2}>{SatisfactionScoreDescription[2]}</Option>
            <Option value={3}>{SatisfactionScoreDescription[3]}</Option>
            <Option value={4}>{SatisfactionScoreDescription[4]}</Option>
            <Option value={5}>{SatisfactionScoreDescription[5]}</Option>
        </Select >
    );
};

export const FoodForm = ({ closeForm, backendUrl }) => {
    const [name, onChangeName] = useState("");
    const [type, onChangeType] = useState("");
    const [desc, onChangeDesc] = useState("");
    const [isLoading, onChangeIsLoading] = useState(false);

    const submitFoodForm = () => {
        const food = {
            Name: name,
            Type: type,
            Description: desc,
        };
        onChangeIsLoading(true);
        fetch(`${backendUrl}/foods`, {
            method: 'POST',
            headers: HEADERS,
            body: JSON.stringify(food)
        })
            .then((resp) => resp.json())
            .then((data) => {
                if (data.message != "Food created") {
                    throw data.message;
                }
                onChangeIsLoading(false);
                closeForm();
            }).catch((error) => {
                Alert.alert("Error", error, [{ text: "Okay" }]);
                onChangeIsLoading(false);
            });
    }

    return (
        <View>
            <Text>Creat Food Stock</Text>
            <Text>Food name</Text>
            <TextInput
                onChangeText={onChangeName}
                value={name}
                placeholder="Enter the food name"
                autoFocus={true}
            />
            <Text>Food type</Text>
            <FoodTypeDropdownList handle={onChangeType} />
            <Text>Food description</Text>
            <TextInput
                onChangeText={onChangeDesc}
                value={desc}
                placeholder="What's about the food?"
            />
            <Button title="Submit" onPress={submitFoodForm} disabled={isLoading} />
            <Button title="Cancel" onPress={closeForm} />
        </View>
    )
}

export const RecordForm = ({ record, closeForm, backendUrl, thisMonthRecord, onChangeThisMonthRecord }) => {

    const eatingDate = record.EatingDate;
    const [name, onChangeName] = useState(record.Food.Name);
    const [quantity, onChangeQuantity] = useState(record.EatenQuantity);
    const [score, onChangeScore] = useState(record.SatisfactionScore);
    const [desc, onChangeDesc] = useState(record.Description);
    const [isLoading, onChangeIsLoading] = useState(false);
    // TODO: handle images uploading
    // const [photoURL, onChangePhotoURL] = useState("");

    const submitRecord = () => {
        const record = {
            Food: { Name: name },
            EatingDate: eatingDate,
            EatenQuantity: Number.parseFloat(quantity),
            SatisfactionScore: score,
            Description: desc
        };
        onChangeIsLoading(true);
        fetch(`${backendUrl}/records`, {
            method: 'POST',
            headers: HEADERS,
            body: JSON.stringify(record)
        })
            .then((resp) => resp.json())
            .then((_) => {
                const d = thisMonthRecord;
                d[record.EatingDate] = record;
                d[record.EatingDate].isModifying = true;
                onChangeThisMonthRecord(d);
                onChangeIsLoading(false);
                closeForm();
            }).catch((error) => {
                Alert.alert("Error", error.message, [{ text: "Okay" }]);
                onChangeIsLoading(false);
            });
    };

    const updateRecord = () => {
        const newRecord = {
            Food: { Name: name },
            EatingDate: eatingDate,
            EatenQuantity: Number.parseFloat(quantity),
            SatisfactionScore: score,
            Description: desc
        };
        const [year, month, day] = eatingDate.split("-");
        onChangeIsLoading(true);
        fetch(`${backendUrl}/records/${year}/${month}/${day}`, {
            method: 'PUT',
            headers: HEADERS,
            body: JSON.stringify(newRecord)
        })
            .then((resp) => resp.json())
            .then((_) => {
                const d = thisMonthRecord;
                d[newRecord.EatingDate] = newRecord;
                d[newRecord.EatingDate].isModifying = true;
                onChangeThisMonthRecord(d);
                onChangeIsLoading(false);
                closeForm();
            }).catch((error) => {
                Alert.alert("Error", error.message, [{ text: "Okay" }]);
                onChangeIsLoading(false);
            });
    }

    return (
        <View>
            <Text>Record today's food</Text>
            <Text>Food name</Text>
            {/* TODO: add auto complete feature from database */}
            <TextInput
                onChangeText={onChangeName}
                value={name}
                defaultValue={name}
                placeholder="Enter the food name"
                autoFocus={true}
            />
            <Text>Satisfaction Score</Text>
            <SatisfactionScoreDropdown initScore={score} handle={onChangeScore} />
            <Text>Eating Date</Text>
            <Text>{eatingDate}</Text>
            <Text>Eaten quantity</Text>
            <TextInput
                onChangeText={onChangeQuantity}
                value={quantity}
                defaultValue={quantity.toString()}
                placeholder="How many bags or cans?"
                keyboardType="numeric"
            />
            <Text>Food description</Text>
            <TextInput
                onChangeText={onChangeDesc}
                value={desc}
                defaultValue={desc}
                placeholder="What's about the food?"
            />
            <Button
                title={record.isModifying ? "Update" : "Add"}
                onPress={record.isModifying ? updateRecord : submitRecord}
                disabled={isLoading}
            />
            <Button title="Cancel" onPress={closeForm} />
        </View>
    )
}
