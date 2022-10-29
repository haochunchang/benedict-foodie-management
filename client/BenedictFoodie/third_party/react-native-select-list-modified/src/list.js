import React, { useRef, useState } from 'react';
import { StyleSheet, View, Modal, TouchableWithoutFeedback, ScrollView, Animated } from 'react-native';


const List = (props) => {
  const AnimatedScrollView = Animated.createAnimatedComponent(ScrollView);

  const [x, onChangeX] = useState(0)
  const [y, onChangeY] = useState(0)
  const [width, onChangeWidth] = useState(0)
  const [height, onChangeHeight] = useState(0)
  const [list, onChangeList] = useState(0)
  const listRef = useRef(null);

  const measureProps = () => {
    onChangeList(height);
    props.select.measureInWindow((x, y, width, height) => {
      onChangeX(x);
      onChangeY(y);
      onChangeWidth(width);
      onChangeHeight(height);
    });
  }

  const { children, position } = props;
  return (
    <Modal transparent={true}>
      <TouchableWithoutFeedback onPress={props.onOverlayPress}>
        <View style={{ flex: 1 }}></View>
      </TouchableWithoutFeedback>
      <View
        onLayout={measureProps}
        ref={listRef}
        style={[
          styles.list,
          {
            width: width,
            maxHeight: props.height,
            left: x,
            top: y + (position === 'down' ? height : -list),
            opacity: list ? 1 : 0,
          },
          props.style
        ]}>
        <View>
          <AnimatedScrollView
            automaticallyAdjustContentInsets={false}
            bounces={false}>
            {
              children.map((item, index) => {
                return (
                  <TouchableWithoutFeedback
                    key={index}
                    onPress={() => { props.onOptionPressed(item.props.value, item.props.children) }}
                  >
                    <View>
                      {item}
                    </View>
                  </TouchableWithoutFeedback>
                );
              })
            }
          </AnimatedScrollView>
        </View>
      </View>
    </Modal>
  );
}

const styles = StyleSheet.create({
  list: {
    position: 'absolute',
    borderWidth: 1,
    borderColor: '#cccccc',
    backgroundColor: 'white',
    borderRadius: 2,
  },
});

module.exports = List;
