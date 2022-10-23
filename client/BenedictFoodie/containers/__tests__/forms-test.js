import React from 'react';
import renderer from 'react-test-renderer';
import { FoodForm, RecordForm } from '../forms';

test('renders FoodForm correctly', () => {
    const tree = renderer.create(<FoodForm />).toJSON();
    expect(tree).toMatchSnapshot();
});

test('renders RecordForm correctly', () => {
    const tree = renderer.create(<RecordForm />).toJSON();
    expect(tree).toMatchSnapshot();
});