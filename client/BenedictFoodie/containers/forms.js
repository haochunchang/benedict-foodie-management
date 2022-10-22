import React from 'react';
import {
    View, Text, TextInput, Button,
} from "react-native"
import { Select, Option } from '../third_party/react-native-select-list-modified';


export const FoodDropdownList = ({ handle }) => {
    return (
        <Select onSelect={(option) => { handle(option.value) }}>
            <Option value={'dry'}>Dry Food</Option>
            <Option value={'wet'}>Wet Food</Option>
            <Option value={'snack'}>Snack</Option>
        </Select >
    );
};

export const FoodForm = ({ closeHandle }) => {
    /**
     * FoodForm consists of
     *  - Name
     *  - Type
     *  - PurchaseDate (default: today's date)
     *  - CurrentQuantity
     *  - Description
     */
    const [name, onChangeName] = React.useState("");
    const [type, onChangeType] = React.useState("");
    const [purchaseDate, onChangePurchaseDate] = React.useState(null);
    const [quantity, onChangeQuantity] = React.useState(0);
    const [desc, onChangeDesc] = React.useState("");

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
            <Button title="Submit" onPress={() => { submitFoodForm(name, type, purchaseDate, quantity, desc) }} />
            <Button title="Cancel" onPress={closeHandle} />
        </View>
    )
}

const submitFoodForm = (name, type, purchaseDate, quantity, desc) => {
    const food = {
        name: name,
        type: type,
        purchaseDate: purchaseDate,
        quantity: quantity,
        description: desc
    }
    console.log(food);
}

export const RecordForm = () => {

}