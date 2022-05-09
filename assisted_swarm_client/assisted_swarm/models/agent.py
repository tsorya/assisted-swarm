# coding: utf-8

"""
    AssistedSwarm

    Assisted swarm  # noqa: E501

    OpenAPI spec version: 1.0.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class Agent(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'id': 'int',
        'status': 'AgentStatus',
        'created_at': 'datetime',
        'terminated_at': 'datetime'
    }

    attribute_map = {
        'id': 'id',
        'status': 'status',
        'created_at': 'created_at',
        'terminated_at': 'terminated_at'
    }

    def __init__(self, id=None, status=None, created_at=None, terminated_at=None):  # noqa: E501
        """Agent - a model defined in Swagger"""  # noqa: E501

        self._id = None
        self._status = None
        self._created_at = None
        self._terminated_at = None
        self.discriminator = None

        if id is not None:
            self.id = id
        if status is not None:
            self.status = status
        if created_at is not None:
            self.created_at = created_at
        if terminated_at is not None:
            self.terminated_at = terminated_at

    @property
    def id(self):
        """Gets the id of this Agent.  # noqa: E501


        :return: The id of this Agent.  # noqa: E501
        :rtype: int
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this Agent.


        :param id: The id of this Agent.  # noqa: E501
        :type: int
        """

        self._id = id

    @property
    def status(self):
        """Gets the status of this Agent.  # noqa: E501


        :return: The status of this Agent.  # noqa: E501
        :rtype: AgentStatus
        """
        return self._status

    @status.setter
    def status(self, status):
        """Sets the status of this Agent.


        :param status: The status of this Agent.  # noqa: E501
        :type: AgentStatus
        """

        self._status = status

    @property
    def created_at(self):
        """Gets the created_at of this Agent.  # noqa: E501


        :return: The created_at of this Agent.  # noqa: E501
        :rtype: datetime
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at):
        """Sets the created_at of this Agent.


        :param created_at: The created_at of this Agent.  # noqa: E501
        :type: datetime
        """

        self._created_at = created_at

    @property
    def terminated_at(self):
        """Gets the terminated_at of this Agent.  # noqa: E501


        :return: The terminated_at of this Agent.  # noqa: E501
        :rtype: datetime
        """
        return self._terminated_at

    @terminated_at.setter
    def terminated_at(self, terminated_at):
        """Sets the terminated_at of this Agent.


        :param terminated_at: The terminated_at of this Agent.  # noqa: E501
        :type: datetime
        """

        self._terminated_at = terminated_at

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(Agent, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, Agent):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other