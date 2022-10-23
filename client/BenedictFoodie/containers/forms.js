import axios from 'axios'
import React, { useState } from 'react';
import {
    View, Text, TextInput, Button, Alert,
} from "react-native"
import { Select, Option } from '../third_party/react-native-select-list-modified';


const FoodDropdownList = ({ handle }) => {
    return (
        <Select onSelect={(value) => { handle(value) }}>
            <Option value={'dry'}>Dry Food</Option>
            <Option value={'wet'}>Wet Food</Option>
            <Option value={'snack'}>Snack</Option>
        </Select >
    );
};

const SatisfactionScoreDropdown = ({ handle }) => {
    return (
        <Select onSelect={(value) => { handle(value) }}>
            <Option value={1}>Does not eat with snack added</Option>
            <Option value={2}>Finish eating with snack</Option>
            <Option value={3}>Finish eating</Option>
            <Option value={4}>Eating eagerly</Option>
            <Option value={5}>Eating eagerly with talking</Option>
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
        // TODO: url should be specified
        sendPostRequest(`${backendUrl}/food`, food, onChangeIsLoading);
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
            <FoodDropdownList handle={onChangeType} />
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

export const RecordForm = ({ today, closeHandle, backendUrl }) => {
    /**
     * RecordForm consists of
     *  - Food name
     *  - Eating Date
     *  - Eaten Quantity
     *  - Satisfaction Score
     *  - Description
     *  - PhotoURL
     */
    const [name, onChangeName] = useState("");
    const [eatingDate, onChangeEatingDate] = useState(today);
    const [quantity, onChangeQuantity] = useState(0);
    const [score, onChangeScore] = useState(1);
    const [desc, onChangeDesc] = useState("");
    const [isLoading, onChangeIsLoading] = useState(false);
    // TODO: handle images uploading
    // const [photoURL, onChangePhotoURL] = useState("");

    const submitRecordForm = () => {
        const record = {
            name: name,
            eatingDate: eatingDate,
            quantity: quantity,
            satisfactionScore: score,
            description: desc
        };
        sendPostRequest(`${backendUrl}/record`, record, onChangeIsLoading);
    };

    return (
        <View>
            <Text>Record today's food</Text>
            <Text>Food name</Text>
            {/* TODO: add auto complete feature from database */}
            <TextInput
                onChangeText={onChangeName}
                value={name}
                placeholder="Enter the food name"
                autoFocus={true}
            />
            <Text>Satisfaction Score</Text>
            <SatisfactionScoreDropdown handle={onChangeScore} />
            <Text>Eating Date</Text>
            <TextInput
                onChangeText={onChangeEatingDate}
                value={eatingDate}
                defaultValue={today}
                keyboardType="numeric"
            />
            <Text>Eaten quantity</Text>
            <TextInput
                onChangeText={onChangeQuantity}
                value={quantity}
                placeholder="How many bags or cans?"
                keyboardType="numeric"
            />
            <Text>Food description</Text>
            <TextInput
                onChangeText={onChangeDesc}
                value={desc}
                placeholder="What's about the food?"
            />
            <Button title="Submit" onPress={submitRecordForm} disabled={isLoading} />
            <Button title="Cancel" onPress={closeHandle} />
        </View>
    )
}

const sendPostRequest = (endpoint, data, onChangeIsLoading) => {
    axios.post(endpoint, data).then((resp) => {
        if (resp.status == 201) {
            Alert.alert("Success", resp.data.message, [{ text: "OK" }]);
            onChangeIsLoading(false);
        } else {
            Alert.alert("Error", resp.data.message, [{ text: "Okay" }]);
        }
    }).catch((error) => {
        Alert.alert("Error", error.message, [{ text: "Okay" }]);
        onChangeIsLoading(false);
    });
}
