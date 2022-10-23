import React, { Component } from 'react';
import { StyleSheet, View, Text, Dimensions, Platform, TouchableWithoutFeedback } from 'react-native';

import List from './list';


class Select extends Component {
  constructor(props) {
    super(props);

    let defaultValue = props.default;

    if (!defaultValue) {
      defaultValue = props.children[0].props.children;
    }

    this.onOptionPressed = this.onOptionPressed.bind(this);

    this.state = {
      value: defaultValue,
      visible: 0,
      listHeight: 0,
    }
  }

  // measureProps() {
  //   this.refs.select.measureInWindow((x, y, width, height) => {
  //     this.setState({
  //       select: {
  //         x: x,
  //         y: y,
  //         width: width,
  //         height: height,
  //       }
  //     });
  //   });
  // }

  onOptionPressed(value, text) {
    this.setState({
      visible: false,
      value: text,
    });

    this.props.onSelect(value, text);
  }

  toggleVisibility() {
    this.setState({ visible: this.state.visible ? 0 : 1 });
  }

  render() {
    const { padding, caret } = this.props;

    let offset = 2 * padding + 15;

    return (
      <View
        style={[styles.container, { zIndex: this.props.zIndex + this.state.visible, paddingBottom: this.state.visible && Platform.OS === 'android' ? this.props.listHeight : 0 }]}>
        <TouchableWithoutFeedback
          onPress={this.toggleVisibility.bind(this)}
        >
          <View
            ref={(view) => { this.select = view; }}
            style={[styles.select, this.props.selectStyle, { paddingHorizontal: padding }]}
          >
            <Text
              style={[this.props.selectTextStyle, { width: 222 - offset }]}
              numberOfLines={1}
              lineBreakMode='tail'
            >{this.state.value}</Text>
          </View>
        </TouchableWithoutFeedback>
        {
          this.state.visible ?
            <List
              style={this.props.listStyle}
              select={this.select}
              height={this.props.listHeight}
              position={this.props.listPosition}
              onOverlayPress={this.toggleVisibility.bind(this)}
              onOptionPressed={this.onOptionPressed}
            >
              {this.props.children}
            </List>
            :
            null
        }
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    alignSelf: 'stretch',
  },
  select: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
    alignSelf: 'stretch',
    height: 40,
    backgroundColor: '#efefef',
    borderRadius: 2,
  },
  labelContainer: {
    justifyContent: 'center',
    height: 25,
    backgroundColor: 'transparent',
  },
});

// Select.propTypes = {
//   listHeight: React.PropTypes.number,
//   listPosition: React.PropTypes.oneOf(['down', 'up']),
//   padding: React.PropTypes.number,
//   zIndex: React.PropTypes.number,
//   onSelect: React.PropTypes.func
// };

Select.defaultProps = {
  listHeight: 100,
  listPosition: 'down',
  padding: 10,
  zIndex: 999,
  onSelect: () => { }
};

module.exports = Select;