import React, { useState } from 'react';
import {
    View, Text, TextInput, Button, Alert,
} from "react-native"
import { Select, Option } from '../third_party/react-native-select-list-modified/src';
import axios from 'axios';

const SatisfactionScoreDescription = {
    1: "Does not eat with snack added",
    2: "Finish eating with snack",
    3: "Finish eating",
    4: "Eating eagerly",
    5: "Eating eagerly with talking",
}

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

export const FoodForm = ({ closeHandle, backendUrl }) => {
    /**
     * FoodForm consists of
     *  - Name
     *  - Type
     *  - PurchaseDate (default: today's date)
     *  - CurrentQuantity
     *  - Description
     */
    const [name, onChangeName] = useState("");
    const [type, onChangeType] = useState("");
    const [quantity, onChangeQuantity] = useState(0);
    const [desc, onChangeDesc] = useState("");
    const [isLoading, onChangeIsLoading] = useState(false);
    const now = new Date();
    const [purchaseDate, onChangePurchaseDate] = useState(`${now.getFullYear()}-${now.getMonth() + 1}-${now.getDate()}`);

    const submitFoodForm = () => {
        const food = {
            Name: name,
            Type: type,
            PurchaseDate: purchaseDate,
            Description: desc,
            Quantity: quantity,
        };
        onChangeIsLoading(true);
        axios.post(`${backendUrl}/foods`, food)
            .then((resp) => {
                if (resp.status == 201 || resp.status == 200) {
                    onChangeIsLoading(false);
                    closeHandle();
                } else {
                    Alert.alert("Error", resp.data.message, [{ text: "Okay" }]);
                }
            }).catch((error) => {
                Alert.alert("Error", error.message, [{ text: "Okay" }]);
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
            <Text>Purchase Date</Text>
            <TextInput
                onChangeText={onChangePurchaseDate}
                value={purchaseDate}
                placeholder="Enter the purchase date in YYYY-MM-DD format"
                keyboardType="numeric"
            />
            <Text>Food quantity</Text>
            <TextInput
                onChangeText={onChangeQuantity}
                value={quantity}
                placeholder="Enter the number of bags or cans"
                keyboardType="phone-pad"
            />
            <Text>Food description</Text>
            <TextInput
                onChangeText={onChangeDesc}
                value={desc}
                placeholder="What's about the food?"
            />
            <Button title="Submit" onPress={submitFoodForm} disabled={isLoading} />
            <Button title="Cancel" onPress={closeHandle} />
        </View>
    )
}

export const RecordForm = ({ record, closeForm, backendUrl, dates, onChangeDates }) => {
    /**
     * RecordForm consists of
     *  - Food name
     *  - Eating Date
     *  - Eaten Quantity
     *  - Satisfaction Score
     *  - Description
     *  - PhotoURL
     */
    const [name, onChangeName] = useState(record.Name);
    const [eatingDate, onChangeEatingDate] = useState(record.EatingDate);
    const [quantity, onChangeQuantity] = useState(
        record.EatenQuantity !== undefined ? record.EatenQuantity.toString() : ""
    );
    const [score, onChangeScore] = useState(record.SatisfactionScore);
    const [desc, onChangeDesc] = useState(record.Description);
    const [isLoading, onChangeIsLoading] = useState(false);
    // TODO: handle images uploading
    // const [photoURL, onChangePhotoURL] = useState("");

    const submitRecord = () => {
        const record = {
            name: name,
            eatingDate: eatingDate,
            quantity: Number.parseFloat(quantity),
            satisfactionScore: score,
            description: desc
        };
        onChangeIsLoading(true);
        axios.post(`${backendUrl}/records`, record)
            .then((resp) => {
                if (resp.status == 201) {
                    const d = dates;
                    d[record.eatingDate] = record;
                    d[record.eatingDate].isModifying = true;
                    onChangeDates(d);
                    onChangeIsLoading(false);
                    closeForm();
                } else {
                    Alert.alert("Error", resp.data.message, [{ text: "Okay" }]);
                }
            }).catch((error) => {
                Alert.alert("Error", error.message, [{ text: "Okay" }]);
                onChangeIsLoading(false);
            });
    };

    const updateRecord = () => {
        const record = {
            name: name,
            eatingDate: eatingDate,
            quantity: Number.parseFloat(quantity),
            satisfactionScore: score,
            description: desc
        };
        onChangeIsLoading(true);
        axios.put(`${backendUrl}/records`, record)
            .then((resp) => {
                if (resp.status == 200) {
                    onChangeIsLoading(false);
                    closeForm();
                } else {
                    Alert.alert("Error", resp.data.message, [{ text: "Okay" }]);
                }
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
            <TextInput
                onChangeText={onChangeEatingDate}
                value={eatingDate}
                defaultValue={eatingDate}
                keyboardType="numeric"
            />
            <Text>Eaten quantity</Text>
            <TextInput
                onChangeText={onChangeQuantity}
                value={quantity}
                defaultValue={quantity}
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
